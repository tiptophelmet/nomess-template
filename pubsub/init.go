package pubsub

import (
	"sync"

	"github.com/tiptophelmet/nomess/config"
	"github.com/tiptophelmet/nomess/logger"
	"github.com/tiptophelmet/nomess/pubsub/broker"
)

type PubSubClient struct {
	broker broker.PubSubBroker
	mu     sync.Mutex
}

var pubsubClient *PubSubClient

func Init() {
	driverConfig := config.Get("pubsub.driver").Required().Str()

	switch driverConfig {
	case "redis":
		pubsubClient = &PubSubClient{broker: &broker.RedisBroker{}}
	case "nats":
		pubsubClient = &PubSubClient{broker: &broker.NATSBroker{}}
	default:
		logger.Panic("unsupported pubsub.driver: %v", driverConfig)
	}

	driverURL := config.Get("pubsub.url").Required().Str()
	pubsubClient.broker.Connect(driverURL)
}

func Connection() broker.PubSubBroker {
	pubsubClient.mu.Lock()
	defer pubsubClient.mu.Unlock()

	return pubsubClient.broker
}
