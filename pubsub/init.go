package pubsub

import (
	"fmt"

	"github.com/tiptophelmet/nomess/config"
	"github.com/tiptophelmet/nomess/logger"
	"github.com/tiptophelmet/nomess/pubsub/broker"
)

var pubsubClient broker.PubSubBroker

func InitPubSub() {
	driverConfig := config.Get("pubsub.driver").Required().Str()

	switch driverConfig {
	case "redis":
		pubsubClient = &broker.RedisBroker{}
	case "nats":
		pubsubClient = &broker.NATSBroker{}
	case "zeromq":
		pubsubClient = &broker.ZeroMQBroker{}
	case "mqtt":
		pubsubClient = &broker.MQTTBroker{}
	default:
		logger.Emergency(fmt.Sprintf("unsupported cache.driver: %v", driverConfig))
	}

	driverURL := config.Get("pubsub.url").Required().Str()
	pubsubClient.Connect(driverURL)
}

func Connection() broker.PubSubBroker {
	return pubsubClient
}
