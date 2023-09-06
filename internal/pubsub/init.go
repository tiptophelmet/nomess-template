package pubsub

import (
	"sync"

	"github.com/tiptophelmet/nomess/internal/config"
	"github.com/tiptophelmet/nomess/internal/logger"
	"github.com/tiptophelmet/nomess/internal/pubsub/broker"
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
