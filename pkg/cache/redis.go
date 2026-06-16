package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mangk/adminBox/pkg/config"
	"github.com/mangk/adminBox/pkg/log"
	"golang.org/x/sync/singleflight"
)

var (
	_redisList      map[string]*redis.Client
	_redisInitOnce  sync.Once
	_redisCacheLock sync.Map
	_redisSG        singleflight.Group
)

func Redis(driver ...string) *redis.Client {
	_redisInitOnce.Do(func() {
		// 初始化redis
		_redisList = make(map[string]*redis.Client)
		for name, redisCfg := range config.CacheCfg() {
			client := redis.NewClient(&redis.Options{
				Addr:     redisCfg.Addr,
				Password: redisCfg.Password, // no password set
				DB:       redisCfg.DB,       // use default DB
			})
			_, err := client.Ping(context.Background()).Result()
			if err != nil {
				log.Panic("redis init error", "name", name, "err", err)
			}
			_redisList[name] = client
		}
	})

	d := "default"
	if len(driver) == 1 {
		d = driver[0]
	}

	if cache, ok := _redisList[d]; ok {
		return cache
	}

	log.Panic("redis driver undefind", "driver", driver)
	return nil
}

func RedisStrGet(key string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()
	redisValue := Redis().Get(ctx, key).Val()

	return redisValue
}

func RedisStrSet(key, value string, exp time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()
	Redis().Set(ctx, key, value, exp)
}

func RedisDel(key string) {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()
	Redis().Del(ctx, key)
}

// Deprecated: Use RedisHOQ() instead.
func RedisHasOrQuery(key string, queryFunc func() (data string, exp time.Duration)) string {
	data := RedisStrGet(key)
	if data == "" {
		d, exp := queryFunc()
		RedisStrSet(key, d, exp)
		data = d
	}

	return data
}

func RedisHasOrQueryByte(key string, queryFunc func() (data []byte, exp time.Duration)) []byte {
	data := RedisStrGet(key)
	if data == "" {
		d, exp := queryFunc()
		RedisStrSet(key, string(d), exp)
		return d
	}

	return []byte(data)
}

// Deprecated: Use RedisHOQGob() instead.
func RedisHasOrQuerySerializerGob[T any](key string, resultReceiver *T, queryFunc func(*T) (expirationTime time.Duration, err error)) error {
	ctx := context.Background()

	// 获取锁
	lock, _ := _redisCacheLock.LoadOrStore(key, &sync.RWMutex{})
	rwLock := lock.(*sync.RWMutex)

	// 尝试读取缓存（读锁）
	rwLock.RLock()
	data, err := Redis().Get(ctx, key).Bytes()
	rwLock.RUnlock()

	if err == nil { // 缓存命中
		// 反序列化并返回
		buffer := bytes.NewBuffer(data)
		decoder := gob.NewDecoder(buffer)
		return decoder.Decode(resultReceiver)
	}

	if err != redis.Nil { // Redis 非缓存缺失错误
		return err
	}

	// 缓存未命中，获取写锁
	rwLock.Lock()
	defer func() {
		rwLock.Unlock()
		// 写锁释放后，清理锁对象
		_redisCacheLock.Delete(key)
	}()

	// 再次检查缓存，防止重复执行查询函数
	data, err = Redis().Get(ctx, key).Bytes()
	if err == nil {
		buffer := bytes.NewBuffer(data)
		decoder := gob.NewDecoder(buffer)
		return decoder.Decode(resultReceiver)
	}

	if err != redis.Nil { // Redis 获取缓存失败
		return err
	}

	// 缓存依然未命中，执行查询函数
	exp, err := queryFunc(resultReceiver)
	if err != nil {
		return err
	}

	// 序列化结果并存储到 Redis
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(*resultReceiver); err != nil {
		return err
	}

	return Redis().Set(ctx, key, buffer.Bytes(), exp).Err()
}

func RedisHOQGob[T any](key string, queryFunc func() (T, time.Duration, error)) (T, error) {
	var zero T
	ctx := context.Background()

	// 先尝试读缓存（无锁）
	data, err := Redis().Get(ctx, key).Bytes()
	if err == nil {
		err = gob.NewDecoder(bytes.NewBuffer(data)).Decode(&zero)
		return zero, err
	}
	if err != redis.Nil {
		return zero, err
	}

	// 缓存未命中：singleflight 保证同一 key 只有一个查询执行
	ret, err, _ := _redisSG.Do(key, func() (interface{}, error) {
		// 二次检查：可能之前的 singleflight 已经写入
		data, err := Redis().Get(ctx, key).Bytes()
		if err == nil {
			return data, nil
		}
		if err != redis.Nil {
			return nil, err
		}

		// 执行查询
		t, exp, err := queryFunc()
		if err != nil {
			return nil, err
		}

		// 序列化
		var buf bytes.Buffer
		if err := gob.NewEncoder(&buf).Encode(t); err != nil {
			return nil, err
		}
		return buf.Bytes(), Redis().Set(ctx, key, buf.Bytes(), exp).Err()
	})
	if err != nil {
		return zero, err
	}

	bytesData, ok := ret.([]byte)
	if !ok {
		return zero, errors.New("redis cache: unexpected type in singleflight result")
	}

	err = gob.NewDecoder(bytes.NewBuffer(bytesData)).Decode(&zero)
	return zero, err
}

// RedisHOQ is a generic function that attempts to retrieve data from Redis cache using the provided key. If the data is not found in the cache, it executes the provided query function to fetch the data, caches it, and returns it. It uses singleflight to ensure that only one query is executed for a given key at a time.
func RedisHOQ[T any](key string, queryFunc func() (data T, exp time.Duration, err error)) (T, error) {
	var zero T
	ctx := context.Background()

	data, err := Redis().Get(ctx, key).Bytes()
	if err == nil {
		var val T
		return val, json.Unmarshal(data, &val)
	}
	if err != redis.Nil {
		return zero, err
	}

	ret, err, _ := _redisSG.Do(key, func() (any, error) {
		data, err := Redis().Get(ctx, key).Bytes()
		if err == nil {
			return data, nil
		}
		if err != redis.Nil {
			return nil, err
		}

		t, exp, err := queryFunc()
		if err != nil {
			return nil, err
		}

		data, err = json.Marshal(t)
		if err != nil {
			return nil, err
		}

		if err := Redis().Set(ctx, key, data, exp).Err(); err != nil {
			return nil, err
		}

		return data, nil
	})
	if err != nil {
		return zero, err
	}

	data, ok := ret.([]byte)
	if !ok {
		return zero, errors.New("redis cache: unexpected type in singleflight result")
	}

	var val T
	return val, json.Unmarshal(data, &val)
}
