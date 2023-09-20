package app

import (
	"os"

	"github.com/tiptophelmet/nomess-core/v5/logger"
	"github.com/tiptophelmet/nomess-template/app/logformat"
)

func initLogger() {
	formatter := &logformat.RequestIdFormatter{}

	if envLevel, present := os.LookupEnv("NOMESS_LOG_LEVEL"); present {
		logger.Init(envLevel, formatter)
	} else {
		logger.Init("error", formatter)
	}
}
