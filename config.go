package main

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
)

type ReviewAppConfig struct {
	BaseApp  string
	Pool     string
	Env      Env
	Services map[string]Service
}

func ParseConfig(path string) (*ReviewAppConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %v", err)
	}
	defer f.Close()

	var config ReviewAppConfig
	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config file: %v", err)
	}

	for name, envVar := range config.Env {
		envVar.Name = name
		config.Env[name] = envVar
	}

	return &config, nil
}
