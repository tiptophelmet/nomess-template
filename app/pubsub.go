package app

import (
	"github.com/tiptophelmet/nomess-core/config"
	"github.com/tiptophelmet/nomess-core/pubsub"
)

func initPubSub() {
	driver := config.Get("pubsub.driver").Required().Str()
	url := config.Get("pubsub.url").Required().Str()

	pubsub.Init(driver, url)
}
