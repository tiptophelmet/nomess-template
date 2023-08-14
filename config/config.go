package config

import (
	"fmt"
	"os"

	"github.com/tiptophelmet/nomess/logger"
	"github.com/tiptophelmet/nomess/util"

	"github.com/pelletier/go-toml"
)

type FallbackConfigs struct {
	list map[string]interface{}
}

func fallback(configName string) interface{} {
	if val, present := fallbackConfigs.list[configName]; present {
		return val
	} else {
		return nil
	}
}

var fallbackConfigs *FallbackConfigs

func initFallbackConfigs() *FallbackConfigs {
	if fallbackConfigs != nil {
		return fallbackConfigs
	}

	var list *toml.Tree

	if tomlData, err := os.ReadFile("../config.toml"); err != nil {
		logger.Alert(err.Error())
	} else if list, err = toml.Load(string(tomlData)); err != nil {
		logger.Alert(err.Error())
	} else {
		illegal := util.GetNonIntersecting(getSupportedConfigKeys(), list.Keys())

		if len(illegal) > 0 {
			logger.Alert(fmt.Sprintf("fallback config.toml has illegal keys: %v", illegal))
		}

	}

	fallbackConfigs = &FallbackConfigs{list.ToMap()}
	return fallbackConfigs
}

type Configs struct {
	list map[string]*Env
}

var appConfigs *Configs

func initAppConfigs() *Configs {
	if appConfigs != nil {
		return appConfigs
	}

	list := make(map[string]*Env)

	for configName, envName := range supportedConfigs {
		list[configName] = initEnv(envName, fallback(configName))
	}

	appConfigs = &Configs{list}
	return appConfigs
}

func Init() {
	initFallbackConfigs()
	initAppConfigs()
}
