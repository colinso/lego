package generator

import (
	"fmt"
	"lego/actions/config"
	"lego/actions/utils"
	"lego/configmodels"
	"os"
)

type CodeGenerator struct {
	cfg configmodels.GeneratorConfig
}

func NewCodeGenerator() *CodeGenerator {
	return &CodeGenerator{
		cfg: config.GetConfig(),
	}
}

type FileGeneratorData struct {
	Name            string
	TemplatePath    string
	FileDestination utils.FileLocation
}

var fileData = []FileGeneratorData{
	{
		Name:            "main",
		TemplatePath:    "./templates/main.tmpl",
		FileDestination: utils.Root,
	},
	{
		Name:            "wire",
		TemplatePath:    "./templates/wire.tmpl",
		FileDestination: utils.Wire,
	},
	{
		Name:            "server",
		TemplatePath:    "./templates/api/server.tmpl",
		FileDestination: utils.API,
	},
}

func (c CodeGenerator) Generate() error {
	// TODO: Run these in parallel
	for _, d := range fileData {
		c.generateBaseConfigFile(d)
	}
	NewHandlerGenerator().Generate()
	NewModelGenerator().Generate()
	NewConfigGenerator().Generate()
	// NewLogicGenerator().Generate()

	fmt.Println(c.cfg)
	return nil
}

func (c CodeGenerator) generateBaseConfigFile(data FileGeneratorData) error {
	tmpl, err := os.ReadFile(data.TemplatePath)
	if err != nil {
		return err
	}
	generateFile(data.Name, string(tmpl), c.cfg, data.FileDestination)
	return nil
}
