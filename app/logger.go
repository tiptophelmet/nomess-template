package app

import (
	"os"

	"github.com/tiptophelmet/nomess-core/v3/logger"
)

func initLogger() {
	if envLevel, present := os.LookupEnv("NOMESS_LOG_LEVEL"); present {
		logger.Init(envLevel)
	} else {
		logger.Init("error")
	}
}
