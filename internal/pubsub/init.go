package pubsub

import (
	"sync"

	"github.com/tiptophelmet/nomess/internal/logger"
	"github.com/tiptophelmet/nomess/internal/pubsub/broker"
)

type pubSubClient struct {
	broker broker.PubSubBroker
	mu     sync.Mutex
}

var client *pubSubClient

func Init(driver, url string) {
	switch driver {
	case "redis":
		client = &pubSubClient{broker: &broker.RedisBroker{}}
	case "nats":
		client = &pubSubClient{broker: &broker.NATSBroker{}}
	default:
		logger.Panic("unsupported pubsub.driver: %v", driver)
	}

	client.broker.Connect(url)
}

func Connection() broker.PubSubBroker {
	client.mu.Lock()
	defer client.mu.Unlock()

	return client.broker
}
