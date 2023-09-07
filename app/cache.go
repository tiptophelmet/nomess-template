package app

import (
	"github.com/tiptophelmet/nomess-template/internal/cache"
	"github.com/tiptophelmet/nomess-template/internal/config"
)

func initCache() {
	driver := config.Get("cache.driver").Required().Str()
	url := config.Get("cache.url").Required().Str()

	cache.Init(driver, url)
}
