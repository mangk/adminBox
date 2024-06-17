package cache

import (
	"context"
	"time"

	"github.com/mangk/adminX/config"
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

func RedisHasOrQuery(key string, queryFunc func() string, exp time.Duration) string {
	data := RedisStrGet(key)
	if data == "" {
		data = queryFunc()
		RedisStrSet(key, data, exp)
	}

	return data
}

type Base64CaptchaStore struct {
}

func (b Base64CaptchaStore) Set(id string, value string) error {
	RedisStrSet("Base64CaptchaStore:"+id, value, time.Duration(config.CaptchaCfg().Overtime)*time.Second)
	return nil
}

func (b Base64CaptchaStore) Get(id string, clear bool) string {
	v := RedisStrGet(id)
	if clear {
		RedisDel(id)
	}
	return v
}

func (b Base64CaptchaStore) Verify(id string, answer string, clear bool) bool {
	v := b.Get("Base64CaptchaStore:"+id, clear)
	return v == answer
}
