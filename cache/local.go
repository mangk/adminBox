package cache

import (
	"time"

	"github.com/mangk/adminBox/config"
	"github.com/songzhibin97/gkit/cache/local_cache"
)

var _blackCacheInit bool
var _blackCache local_cache.Cache

func Local() local_cache.Cache {
	if !_blackCacheInit {
		_blackCache = local_cache.NewCache()
		_blackCacheInit = true
	}

	return _blackCache
}

// TODO 验证码使用本地缓存，在多机部署时会存在问题
type Base64CaptchaStore struct {
}

func (b Base64CaptchaStore) Set(id string, value string) error {
	_blackCache.Set("Base64CaptchaStore:"+id, value, time.Duration(config.CaptchaCfg().Overtime)*time.Second)
	return nil
}

func (b Base64CaptchaStore) Get(id string, clear bool) string {
	v, has := _blackCache.Get("Base64CaptchaStore:" + id)
	if has {
		if clear {
			_blackCache.Delete("Base64CaptchaStore:" + id)
		}
		return v.(string)
	}
	return ""
}

func (b Base64CaptchaStore) Verify(id string, answer string, clear bool) bool {
	v := b.Get(id, clear)
	return v == answer
}
