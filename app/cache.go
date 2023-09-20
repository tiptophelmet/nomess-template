package app

import (
	"github.com/tiptophelmet/nomess-core/v5/cache"
	"github.com/tiptophelmet/nomess-core/v5/config"
)

func initCache() {
	driver := config.Get("cache.driver").Required().Str()
	url := config.Get("cache.url").Required().Str()

	cache.Init(driver, url)
}
