package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/log"
)

var (
	_redisList     map[string]*redis.Client
	_redisInitOnce sync.Once
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

func RedisHasOrQuery(key string, queryFunc func() (data string, exp time.Duration)) string {
	data := RedisStrGet(key)
	if data == "" {
		d, exp := queryFunc()
		RedisStrSet(key, d, exp)
		data = d
	}

	return data
}

func RedisHasOrQuerySerializerGob[T any](key string, resultReceiver *T, queryFunc func(*T) (expirationTime time.Duration, err error)) error {
	data, err := Redis().Get(context.Background(), key).Bytes()
	if err != nil && err != redis.Nil {
		return err
	}

	if err == redis.Nil {
		exp, err := queryFunc(resultReceiver)
		if err != nil {
			return err
		}
		var buffer bytes.Buffer
		encoder := gob.NewEncoder(&buffer)
		if err := encoder.Encode(*resultReceiver); err != nil {
			return err
		}

		Redis().Set(context.Background(), key, buffer.Bytes(), exp)
	} else {
		buffer := bytes.NewBuffer(data)
		decoder := gob.NewDecoder(buffer)
		if err := decoder.Decode(resultReceiver); err != nil {
			return err
		}
	}
	return nil
}
