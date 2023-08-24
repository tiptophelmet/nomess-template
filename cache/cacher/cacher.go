package cacher

import "time"

type Cacher interface {
	Connect(url string) error

	Set(key string, val []byte, namespace string, exp time.Duration) error
	Has(key string, namespace string) bool
	Get(key string, namespace string) ([]byte, error)

	Expire(exp time.Duration)
	ExpireTime() time.Duration

	Delete(key string, namespace string) error
	Flush() error
}
