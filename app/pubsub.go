package app

import (
	"github.com/tiptophelmet/nomess-template/internal/config"
	"github.com/tiptophelmet/nomess-template/internal/pubsub"
)

func initPubSub() {
	driver := config.Get("pubsub.driver").Required().Str()
	url := config.Get("pubsub.url").Required().Str()

	pubsub.Init(driver, url)
}
