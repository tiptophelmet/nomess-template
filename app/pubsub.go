package app

import (
	"github.com/tiptophelmet/nomess-core/v4/config"
	"github.com/tiptophelmet/nomess-core/v4/pubsub"
)

func initPubSub() {
	driver := config.Get("pubsub.driver").Required().Str()
	url := config.Get("pubsub.url").Required().Str()

	pubsub.Init(driver, url)
}
