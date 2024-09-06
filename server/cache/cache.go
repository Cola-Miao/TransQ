package cache

import (
	. "github.com/Cola-Miao/TransQ/server/config"
	"github.com/Cola-Miao/TransQ/server/format"
	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

func InitCache() error {
	format.FuncStart("InitCache")
	defer format.FuncEnd("InitCache")

	c := cache.New(Cfg.Cache.DefaultExpiration, Cfg.Cache.CleanupInterval)
	Cache = c

	return nil
}
