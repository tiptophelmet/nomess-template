package cache

import (
	"sync"

	"github.com/tiptophelmet/nomess/cache/cacher"
	"github.com/tiptophelmet/nomess/config"
	"github.com/tiptophelmet/nomess/logger"
)

type CacheClient struct {
	cacher cacher.Cacher
	mu     sync.Mutex
}

var cacheClient *CacheClient

func Init() {
	if cacheClient != nil {
		return
	}

	driverConfig := config.Get("cache.driver").Required().Str()

	switch driverConfig {
	case "redis":
		cacheClient = &CacheClient{cacher: cacher.InitRedisCacher()}
	case "memcached":
		cacheClient = &CacheClient{cacher: cacher.InitMemcachedCacher()}
	case "aerospike":
		cacheClient = &CacheClient{cacher: cacher.InitAerospikeCacher()}
	default:
		logger.Panic("unsupported cache.driver: %v", driverConfig)
	}

	driverURL := config.Get("cache.url").Required().Str()
	cacheClient.cacher.Connect(driverURL)
}

func Connection() cacher.Cacher {
	cacheClient.mu.Lock()
	defer cacheClient.mu.Unlock()

	return cacheClient.cacher
}
