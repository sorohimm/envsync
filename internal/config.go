package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Replace map[string]string

func NewConfig() *Config {
	return &Config{}
}

type Config struct {
	Replace Replace `yaml:"replace"`
}

func loadConfig(path string) (*Config, error) {
	configFile, err := os.ReadFile(path)
	if err != nil {
		return NewConfig(), fmt.Errorf("Read config file error: %w\n", err)
	}

	var cfg Config
	err = yaml.Unmarshal(configFile, &cfg)
	if err != nil {
		return NewConfig(), fmt.Errorf("Unmarshal config YAML error: %w\n", err)
	}

	return &cfg, nil
}
