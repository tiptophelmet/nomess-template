package config

import (
	"os"

	"github.com/tiptophelmet/nomess/logger"
)

func initEnv(name string, fallback interface{}) *Env {
	var value interface{}

	if envValue, present := os.LookupEnv(name); present {
		value = envValue
	} else {
		value = nil
	}

	if value == nil && fallback == nil {
		logger.Fatal("env %s is not set & has no fallback", name)
	}

	return &Env{name, value, fallback}
}

type Env struct {
	name     string
	value    interface{}
	fallback interface{}
}
