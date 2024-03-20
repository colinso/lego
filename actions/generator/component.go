package generator

import (
	"lego/actions/config"
	"lego/actions/utils"
	"lego/configmodels"
	"os"
)

const componentTemplatePath = "./templates/component.tmpl"

type ComponentGenerator struct {
	cfg configmodels.GeneratorConfig
}

func NewComponentGenerator() *ComponentGenerator {
	return &ComponentGenerator{
		cfg: config.GetConfig(),
	}
}

func (c ComponentGenerator) Generate(packageName string, constructorNames []string, location utils.FileLocation) error {
	componentConfig := struct {
		PackageName string
		Objects     []string
	}{
		PackageName: packageName,
		Objects:     constructorNames,
	}

	componentTemplate, err := os.ReadFile(componentTemplatePath)
	if err != nil {
		return err
	}

	generateFile("component", string(componentTemplate), componentConfig, location)
	return nil
}
