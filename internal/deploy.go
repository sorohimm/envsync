package internal

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Env map[string]string

type Deploy struct {
	Env   Env   `yaml:"env"`
	Vault Vault `yaml:"vault"`
}

type Vault struct {
	Services []Service `yaml:"env"`
}

type Service map[string]Env

// loadDeployEnv reads a deploy file from the specified path.
func loadDeployEnv(path string) (*Deploy, error) {
	deployFileRaw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Read deploy file error: %w\n", err)
	}

	deployFile := cutDeployFile(deployFileRaw)

	var deploy Deploy
	err = yaml.Unmarshal(deployFile, &deploy)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal deploy YAML error: %w\n", err)
	}

	return &deploy, nil
}

func cutDeployFile(file []byte) []byte {
	res, _, _ := strings.Cut(string(file), "envoy:")
	return []byte(res)
}
