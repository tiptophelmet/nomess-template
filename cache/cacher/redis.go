package cacher

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tiptophelmet/nomess/errs"
	"github.com/tiptophelmet/nomess/logger"
)

type RedisCacher struct {
	client *redis.Client
}

func InitRedisCacher() *RedisCacher {
	cache := &RedisCacher{}
	return cache
}

func (rc *RedisCacher) Connect(url string) error {
	options, err := redis.ParseURL(url)
	if err != nil {
		logger.Err(fmt.Sprintf("failed to connect to redis: %v", err.Error()))
		return errs.ErrCacheStoreConnectionFailed
	}

	rc.client = redis.NewClient(options)
	return nil
}

func (rc *RedisCacher) Set(key string, val []byte, namespace string, exp time.Duration) error {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	if namespace != "" {
		key = fmt.Sprintf("%v.%v", namespace, key)
	}

	return rc.client.Set(ctx, key, val, exp).Err()
}

func (rc *RedisCacher) Has(key string, namespace string) bool {
	return false
}

func (rc *RedisCacher) Get(key string, namespace string) ([]byte, error) {
	byteSlice := make([]byte, 0)
	return byteSlice, nil
}

func (rc *RedisCacher) Expire(exp time.Duration) {

}

func (rc *RedisCacher) ExpireTime() time.Duration {
	duration, _ := time.ParseDuration("0")
	return duration
}

func (rc *RedisCacher) Delete(key string, namespace string) error {
	return nil
}

func (rc *RedisCacher) Flush() error {
	return nil
}
