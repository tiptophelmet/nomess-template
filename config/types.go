package config

import (
	"fmt"

	"github.com/tiptophelmet/nomess/errs"
	"github.com/tiptophelmet/nomess/logger"
)

func raw(name string) (interface{}, error) {
	env, found := appConfigs.list[name]
	if !found {
		logger.Err(fmt.Sprintf("%s: \"%s\"", errs.ErrConfigNotFound.Error(), name))
		return nil, errs.ErrConfigNotFound
	}

	var rawVal any

	if env.value != nil {
		rawVal = env.value
	} else {
		rawVal = env.fallback
	}

	return rawVal, nil
}

func Str(name string) (string, error) {
	rawVal, err := raw(name)
	if err != nil {
		return "", err
	}

	val, typeOk := rawVal.(string)
	if !typeOk {
		logger.Err(fmt.Sprintf("%s: \"%s\"", err.Error(), name))
		return "", errs.ErrTypeAssertion
	}

	return val, nil
}

func Int(name string) (int, error) {
	rawVal, err := raw(name)
	if err != nil {
		return 0, err
	}

	val, typeOk := rawVal.(int)
	if !typeOk {
		logger.Err(fmt.Sprintf("%s: \"%s\"", err.Error(), name))
		return 0, errs.ErrTypeAssertion
	}

	return val, nil
}

func Float(name string) (float32, error) {
	rawVal, err := raw(name)
	if err != nil {
		return 0.0, err
	}

	val, typeOk := rawVal.(float32)
	if !typeOk {
		logger.Err(fmt.Sprintf("%s: \"%s\"", err.Error(), name))
		return 0.0, errs.ErrTypeAssertion
	}

	return val, nil
}

func Bool(name string) (bool, error) {
	rawVal, err := raw(name)
	if err != nil {
		return false, err
	}

	val, typeOk := rawVal.(bool)
	if !typeOk {
		logger.Err(fmt.Sprintf("%s: \"%s\"", err.Error(), name))
		return false, errs.ErrTypeAssertion
	}

	return val, nil
}
