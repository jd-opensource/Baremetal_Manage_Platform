package global

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	Cache *cache.Cache
)

func init() {
	Cache = cache.New(10*time.Minute, 1*time.Minute)
}

func Set(k string, v interface{}) {
	if Cache == nil {
		return
	}

	Cache.Set(k, v, cache.DefaultExpiration)
}

func Get(k string) (interface{}, bool) {
	if Cache == nil {
		return nil, false
	}
	return Cache.Get(k)
}
