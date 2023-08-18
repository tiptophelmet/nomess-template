package config

import (
	"fmt"

	"github.com/tiptophelmet/nomess/logger"
	"github.com/tiptophelmet/nomess/util"
)

type ConfigOptions struct {
	name   string
	rawVal interface{}
}

func Get(name string) *ConfigOptions {
	options := &ConfigOptions{
		name:   name,
		rawVal: raw(name),
	}

	return options
}

func raw(name string) interface{} {
	env, found := configList.list[name]
	if !found {
		return nil
	}

	var rawVal interface{}

	if env.value != nil {
		rawVal = env.value
	} else {
		rawVal = env.fallback
	}

	return rawVal
}

func (co *ConfigOptions) Required() *ConfigOptions {
	if util.IsEmpty(co.rawVal) {
		logger.Emergency(fmt.Sprintf("could not resolve config %v", co.name))
		return nil
	}

	return co
}

func (co *ConfigOptions) Str() string {
	val, typeOk := co.rawVal.(string)
	if !typeOk {
		logger.Err(fmt.Sprintf("could not assert config %v to string", co.name))
		return ""
	}

	return val
}

func (co *ConfigOptions) Int() int {
	val, typeOk := co.rawVal.(int)
	if !typeOk {
		logger.Err(fmt.Sprintf("could not assert config %v to int", co.name))
		return 0
	}

	return val
}

func (co *ConfigOptions) Int64() int64 {
	val, typeOk := co.rawVal.(int64)
	if !typeOk {
		logger.Err(fmt.Sprintf("could not assert config %v to int64", co.name))
		return 0
	}

	return val
}

func (co *ConfigOptions) Float() float32 {
	val, typeOk := co.rawVal.(float32)
	if !typeOk {
		logger.Err(fmt.Sprintf("could not assert config %v to float", co.name))
		return 0.0
	}

	return val
}

func (co *ConfigOptions) Bool() bool {
	val, typeOk := co.rawVal.(bool)
	if !typeOk {
		logger.Err(fmt.Sprintf("could not assert config %v to bool", co.name))
		return false
	}

	return val
}
