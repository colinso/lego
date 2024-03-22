package config

import (
	"os"

	configmodels "github.com/colinso/lego/config/models"
	"github.com/colinso/lego/utils/datastructures"
	"gopkg.in/yaml.v3"
)

var cfg *configmodels.Base

func ParseConfig(configPath string, projectPath string) {
	file, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	yamlCfg := configmodels.BaseYaml{}
	yaml.Unmarshal(file, &yamlCfg)

	cfg = &configmodels.Base{
		ProjectPath: yamlCfg.ProjectPath,
		Name:        yamlCfg.Name,
		AppConfig:   yamlCfg.AppConfig,
		Models:      yamlCfg.Models,
		HTTP:        yamlCfg.HTTP,
		Database:    yamlCfg.Database,
	}

	services := []configmodels.Service{}
	for _, s := range yamlCfg.Services {
		methods := []configmodels.Method{}
		for _, m := range s.Methods {
			accepts, err := datastructures.NewOrderedMap[string, string]().FromYaml(m.Accepts)
			if err != nil {
				panic(err)
			}
			method := configmodels.Method{
				Name:    m.Name,
				Returns: m.Returns,
				Accepts: *accepts,
			}

			methods = append(methods, method)
		}
		services = append(services, configmodels.Service{
			Name:    s.Name,
			Methods: methods,
		})
	}
	cfg.Services = services

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
