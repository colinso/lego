package config

import (
	"os"

	configmodels "github.com/colinso/lego/config/models"

	"gopkg.in/yaml.v3"
)

var cfg *configmodels.Base

func ParseConfig(configPath string, projectPath string) {
	file, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(file, &cfg)
	cfg.ProjectPath = projectPath
	// Ensure essential configs are in place
	cfg.Validate()
}

// TODO: This is test damage, we don't need to expose this
func SetConfig(config configmodels.Base) {
	cfg = &config
}

func GetConfig() configmodels.Base {
	if cfg == nil {
		panic("No configuration has been parsed")
	}
	return *cfg
}
