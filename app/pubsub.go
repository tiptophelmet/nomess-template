package app

import (
	"github.com/tiptophelmet/nomess/internal/config"
	"github.com/tiptophelmet/nomess/internal/pubsub"
)

func initPubSub() {
	driver := config.Get("pubsub.driver").Required().Str()
	url := config.Get("pubsub.url").Required().Str()

	pubsub.Init(driver, url)
}
