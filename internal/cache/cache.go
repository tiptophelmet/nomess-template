package cache

import (
	"sync"

	"github.com/tiptophelmet/nomess/internal/cache/cacher"
	"github.com/tiptophelmet/nomess/internal/config"
	"github.com/tiptophelmet/nomess/internal/logger"
)

type cacheClient struct {
	cacher cacher.Cacher
	mu     sync.Mutex
}

var client *cacheClient

func Init() {
	if client != nil {
		return
	}

	driverConfig := config.Get("cache.driver").Required().Str()

	switch driverConfig {
	case "redis":
		client = &cacheClient{cacher: cacher.InitRedisCacher()}
	case "memcached":
		client = &cacheClient{cacher: cacher.InitMemcachedCacher()}
	case "aerospike":
		client = &cacheClient{cacher: cacher.InitAerospikeCacher()}
	default:
		logger.Panic("unsupported cache.driver: %v", driverConfig)
	}

	driverURL := config.Get("cache.url").Required().Str()
	client.cacher.Connect(driverURL)
}

func Connection() cacher.Cacher {
	client.mu.Lock()
	defer client.mu.Unlock()

	return client.cacher
}
