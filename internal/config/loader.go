package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

func LoadConfig() (Config, error) {
	var config Config

	err := envconfig.Process("", &config)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return config, err
	}

	return config, nil
}
