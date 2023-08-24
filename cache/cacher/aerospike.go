package cacher

import (
	"fmt"
	"time"

	"github.com/aerospike/aerospike-client-go/v6"
	aerospikeurl "github.com/tiptophelmet/aerospike-url"
	"github.com/tiptophelmet/nomess/errs"
	"github.com/tiptophelmet/nomess/logger"
)

type AerospikeCacher struct {
	client *aerospike.Client
}

func InitAerospikeCacher() *AerospikeCacher {
	cacher := &AerospikeCacher{}
	return cacher
}

func (ac *AerospikeCacher) Connect(url string) error {
	clientFactory, err := aerospikeurl.Parse(url)
	if err != nil {
		logger.Err(fmt.Sprintf("failed to connect to aerospike: %v", err.Error()))
		return errs.ErrCacheStoreConnectionFailed
	}

	ac.client, err = clientFactory.BuildClient()
	if err != nil {
		logger.Err(fmt.Sprintf("failed to connect to aerospike: %v", err.Error()))
		return errs.ErrCacheStoreConnectionFailed
	}

	return nil
}

func (ac *AerospikeCacher) Set(key string, val []byte, namespace string, exp time.Duration) error {
	if namespace != "" {
		key = fmt.Sprintf("%v.%v", namespace, key)
	}

	setKey := fmt.Sprintf("set-%v", key)
	aeroKey, err := aerospike.NewKey(namespace, setKey, key)
	if err != nil {
		return err
	}

	bin := aerospike.NewBin(namespace, val)

	err = ac.client.PutBins(nil, aeroKey, bin)
	if err != nil {
		return err
	}

	return nil
}

func (ac *AerospikeCacher) Has(key string, namespace string) bool {
	return false
}

func (ac *AerospikeCacher) Get(key string, namespace string) ([]byte, error) {
	byteSlice := make([]byte, 0)
	return byteSlice, nil
}

func (ac *AerospikeCacher) Expire(exp time.Duration) {

}

func (ac *AerospikeCacher) ExpireTime() time.Duration {
	duration, _ := time.ParseDuration("0")
	return duration
}

func (ac *AerospikeCacher) Delete(key string, namespace string) error {
	return nil
}

func (ac *AerospikeCacher) Flush() error {
	return nil
}
