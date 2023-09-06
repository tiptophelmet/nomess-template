package pubsub

import (
	"sync"

	"github.com/tiptophelmet/nomess/config"
	"github.com/tiptophelmet/nomess/logger"
	"github.com/tiptophelmet/nomess/pubsub/broker"
)

type pubSubClient struct {
	broker broker.PubSubBroker
	mu     sync.Mutex
}

var client *pubSubClient

func Init() {
	driverConfig := config.Get("pubsub.driver").Required().Str()

	switch driverConfig {
	case "redis":
		client = &pubSubClient{broker: &broker.RedisBroker{}}
	case "nats":
		client = &pubSubClient{broker: &broker.NATSBroker{}}
	default:
		logger.Panic("unsupported pubsub.driver: %v", driverConfig)
	}

	driverURL := config.Get("pubsub.url").Required().Str()
	client.broker.Connect(driverURL)
}

func Connection() broker.PubSubBroker {
	client.mu.Lock()
	defer client.mu.Unlock()

	return client.broker
}
