package config

import (
	"fmt"
	"os"

	"github.com/tiptophelmet/nomess/logger"
	"github.com/tiptophelmet/nomess/util"

	"github.com/BurntSushi/toml"
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

	list := make(map[string]interface{})

	if tomlData, err := os.ReadFile("../config.toml"); err != nil {
		logger.Alert(err.Error())
	} else if md, err := toml.Decode(string(tomlData), &list); err != nil {
		logger.Alert(err.Error())
	} else {
		actual := util.ConvertTomlKeysToStrings(md.Keys())
		illegal := util.GetNonIntersecting(getSupportedConfigKeys(), actual)

		if len(illegal) > 0 {
			logger.Alert(fmt.Sprintf("fallback config.toml has illegal keys: %v", illegal))
		}

	}

	fallbackConfigs = &FallbackConfigs{list}
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
