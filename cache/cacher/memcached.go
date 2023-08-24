package cacher

import (
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/tiptophelmet/nomess/errs"
	"github.com/tiptophelmet/nomess/logger"
)

type MemcachedCacher struct {
	client *memcache.Client
}

func InitMemcachedCacher() *MemcachedCacher {
	cache := &MemcachedCacher{}
	return cache
}

func (mc *MemcachedCacher) Connect(url string) error {
	mc.client = memcache.New(url)

	err := mc.client.Ping()
	if err != nil {
		logger.Err(fmt.Sprintf("failed to connect to memcached: %v", err.Error()))
		return errs.ErrCacheStoreConnectionFailed
	}

	return nil
}

func (mc *MemcachedCacher) Set(key string, val []byte, namespace string, exp time.Duration) error {
	if namespace != "" {
		key = fmt.Sprintf("%v.%v", namespace, key)
	}

	return mc.client.Set(&memcache.Item{
		Key:        key,
		Value:      val,
		Expiration: int32(exp.Seconds()),
	})
}

func (mc *MemcachedCacher) Has(key string, namespace string) bool {
	return false
}

func (mc *MemcachedCacher) Get(key string, namespace string) ([]byte, error) {
	byteSlice := make([]byte, 0)
	return byteSlice, nil
}

func (mc *MemcachedCacher) Expire(exp time.Duration) {

}

func (mc *MemcachedCacher) ExpireTime() time.Duration {
	duration, _ := time.ParseDuration("0")
	return duration
}

func (mc *MemcachedCacher) Delete(key string, namespace string) error {
	return nil
}

func (mc *MemcachedCacher) Flush() error {
	return nil
}