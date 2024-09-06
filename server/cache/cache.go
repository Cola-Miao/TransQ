package cache

import (
	. "github.com/Cola-Miao/TransQ/server/config"
	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

func InitCache() error {
	c := cache.New(Cfg.Cache.DefaultExpiration, Cfg.Cache.CleanupInterval)
	Cache = c

	return nil
}
