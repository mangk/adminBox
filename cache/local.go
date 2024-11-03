package cache

import (
	"time"

	"github.com/mangk/adminBox/config"
)

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
