package app

import (
	"github.com/tiptophelmet/nomess-core/v4/cache"
	"github.com/tiptophelmet/nomess-core/v4/config"
)

func initCache() {
	driver := config.Get("cache.driver").Required().Str()
	url := config.Get("cache.url").Required().Str()

	cache.Init(driver, url)
}
