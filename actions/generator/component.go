package generator

import (
	"os"

	"github.com/colinso/lego/actions/config"
	"github.com/colinso/lego/actions/utils"
	"github.com/colinso/lego/configmodels"
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
