package config

import (
	"encoding/json"
	"os"
)

var globalConfig GlobalConfig

func LoadConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &globalConfig)
}

func GetGlobalConfig() GlobalConfig {
	return globalConfig
}
