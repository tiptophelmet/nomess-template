package config

import (
	"fmt"
	"os"

	"github.com/tiptophelmet/nomess/logger"
	"github.com/tiptophelmet/nomess/util"

	"github.com/pelletier/go-toml"
)

type FallbackConfigList struct {
	list map[string]interface{}
}

func fallback(configName string) interface{} {
	if val, present := fallbackList.list[configName]; present {
		return val
	} else {
		return nil
	}
}

var fallbackList *FallbackConfigList

func initFallbackConfigs() *FallbackConfigList {
	if fallbackList != nil {
		return fallbackList
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

	fallbackList = &FallbackConfigList{list.ToMap()}
	return fallbackList
}

type ConfigList struct {
	list map[string]*Env
}

var configList *ConfigList

func initAppConfigs() *ConfigList {
	if configList != nil {
		return configList
	}

	list := make(map[string]*Env)

	for configName, envName := range supportedConfigs {
		list[configName] = initEnv(envName, fallback(configName))
	}

	configList = &ConfigList{list}
	return configList
}

func Init() {
	initFallbackConfigs()
	initAppConfigs()
}
