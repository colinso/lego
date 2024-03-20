package actions

import (
	"fmt"
	"lego/actions/config"
	"lego/actions/utils"
	"lego/configmodels"
	"os"
	"text/template"
)

const modelTemplatePath = "./templates/model.tmpl"
const handlerTemplatePath = "./templates/handler.tmpl"
const componentTemplatePath = "./templates/component.tmpl"

type CodeGenerator struct {
	cfg configmodels.GeneratorConfig
}

func NewCodeGenerator() *CodeGenerator {
	return &CodeGenerator{
		cfg: config.GetConfig(),
	}
}

func (c CodeGenerator) Generate(path string) error {
	// TODO: Run these in parallel
	c.generateModels()
	c.generateHandlers()

	fmt.Println(c.cfg)
	return nil
}

func (c CodeGenerator) generateModels() error {
	modelTemplate, err := os.ReadFile(modelTemplatePath)
	if err != nil {
		return err
	}

	for _, m := range c.cfg.Models {
		generateFile(m.Name, string(modelTemplate), m, utils.Model)
	}
	return nil
}

func (c CodeGenerator) generateHandlers() error {
	handlerTemplate, err := os.ReadFile(handlerTemplatePath)
	if err != nil {
		return err
	}

	c.generateComponent("handlers", getHandlerNames(c.cfg.HTTP))

	for _, h := range c.cfg.HTTP {
		generateFile(h.Name, string(handlerTemplate), h, utils.Handler)
	}
	return nil
}

func (c CodeGenerator) generateComponent(packageName string, constructorNames []string) error {
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

	generateFile("component", string(componentTemplate), componentConfig, utils.Handler)
	return nil
}

func getHandlerNames(handlers []configmodels.HTTPHandlerConfig) []string {
	handlerNames := make([]string, 0)
	for _, h := range handlers {
		handlerNames = append(handlerNames, h.Name+"Handler")
	}
	return handlerNames
}

func generateFile(name string, templateFile string, config any, ftype utils.FileLocation) error {
	tmpl, err := template.New(name).Funcs(utils.TemplateFuncs).Parse(string(templateFile))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	f := utils.CreateFileForType(ftype, name)

	err = tmpl.Execute(f, config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
	return nil
}
