package app

import (
	"github.com/tiptophelmet/nomess-core/cache"
	"github.com/tiptophelmet/nomess-core/config"
)

func initCache() {
	driver := config.Get("cache.driver").Required().Str()
	url := config.Get("cache.url").Required().Str()

	cache.Init(driver, url)
}
