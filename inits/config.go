package inits

import (
	"airhass/config"
	"os"

	"gopkg.in/yaml.v3"
)

func Config() error {
	// Read config file
	configFileBytes, err := os.ReadFile("config.yml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configFileBytes, &config.Config)
	if err != nil {
		return err
	}

	return nil
}
