package config

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

var (
	globalConfig GlobalConfig
	cfg          Config
)

func LoadConfig(path string) error {
	if strings.HasSuffix(path, ".json") {
		return loadJSONConfig(path)
	}

	if strings.HasSuffix(path, ".toml") {
		return loadTomlConfig(path)
	}

	return errors.New("config type not support")
}

func loadJSONConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &globalConfig)
}

func loadTomlConfig(path string) error {
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return err
	}

	if len(cfg.Priority) == 0 {
		return errors.New("priority list is empty")
	}
	if cfg.Paths.Resources == "" {
		return errors.New("paths.resources is required")
	}

	return nil
}

func GetGlobalConfig() GlobalConfig {
	return globalConfig
}

func GetConfig() Config {
	return cfg
}
