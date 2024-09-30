package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/log"
	"github.com/mangk/adminBox/moduleRegister"
	"github.com/songzhibin97/gkit/cache/local_cache"
)

var _blackCache local_cache.Cache
var _redisList map[string]*redis.Client

func init() {
	moduleRegister.ModuleAdd(cache{})
}

type cache struct{}

func (cache) InitModule() {
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

	// 初始化本地缓存
	_blackCache = local_cache.NewCache()
}

func Redis(driver ...string) *redis.Client {
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

func Local() local_cache.Cache {
	return _blackCache
}
