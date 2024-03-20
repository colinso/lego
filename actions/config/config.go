package config

import (
	"lego/configmodels"
	"os"

	"gopkg.in/yaml.v3"
)

var cfg *configmodels.GeneratorConfig

func ParseConfig(configPath string, projectPath string) {
	file, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(file, &cfg)
	cfg.ProjectPath = projectPath
	// Ensure essential configs are in place
	cfg.AppConfig.Validate()
}

// TODO: This is test damage, we don't need to expose this
func SetConfig(config configmodels.GeneratorConfig) {
	cfg = &config
}

func GetConfig() configmodels.GeneratorConfig {
	if cfg == nil {
		panic("No configuration has been parsed")
	}
	return *cfg
}
