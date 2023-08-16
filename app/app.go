package app

import (
	"github.com/tiptophelmet/nomess/config"
	"github.com/tiptophelmet/nomess/intl"
	"github.com/tiptophelmet/nomess/logger"
	"github.com/tiptophelmet/nomess/router"
)

func Init() {
	logger.Init()

	config.Init()

	intl.Init("en-US")

	router.Init()
}
