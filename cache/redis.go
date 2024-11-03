package cache

import (
	"context"
	"time"

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
