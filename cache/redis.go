package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	"github.com/go-redis/redis/v8"
)

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

func RedisHasOrQuerySerializerGob[T any](key string, resultReceiver *T, queryFunc func(*T) (expirationTime time.Duration, error error)) error {
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
