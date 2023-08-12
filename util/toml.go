package util

import "github.com/BurntSushi/toml"

func ConvertTomlKeysToStrings(tomlKeys []toml.Key) []string {
	var keys []string

	for _, val := range tomlKeys {
		keys = append(keys, val.String())
	}

	return keys
}
