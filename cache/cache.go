package cache

import (
	"fmt"

	"github.com/tiptophelmet/nomess/cache/cacher"
	"github.com/tiptophelmet/nomess/config"
	"github.com/tiptophelmet/nomess/logger"
)

var cacheClient cacher.Cacher

func InitCache() {
	driverConfig := config.Get("cache.driver").Required().Str()

	switch driverConfig {
	case "redis":
		cacheClient = cacher.InitRedisCacher()
	case "memcached":
		cacheClient = cacher.InitMemcachedCacher()
	case "aerospike":
		cacheClient = cacher.InitAerospikeCacher()
	default:
		logger.Emergency(fmt.Sprintf("unsupported cache.driver: %v", driverConfig))
	}

	driverURL := config.Get("cache.url").Required().Str()
	cacheClient.Connect(driverURL)
}

func Connection() cacher.Cacher {
	return cacheClient
}
