package generator

import (
	"github.com/colinso/lego/actions/config"
	"github.com/colinso/lego/actions/utils"
	"github.com/colinso/lego/configmodels"
)

type CodeGenerator struct {
	cfg configmodels.GeneratorConfig
}

func NewCodeGenerator() *CodeGenerator {
	return &CodeGenerator{
		cfg: config.GetConfig(),
	}
}

type FileDataParams struct {
	Name            string
	Extension       utils.FileExtension
	TemplatePath    string
	FileDestination utils.FileLocation
	ConfigModel     any
	HasConstructor  bool
}

func (c CodeGenerator) Generate() error {
	config := c.cfg
	// Generate individual files
	filesToGenerate := []FileDataParams{
		{
			Name:            "main",
			TemplatePath:    "./templates/main.tmpl",
			FileDestination: utils.Cmd,
			ConfigModel:     config,
			Extension:       utils.Go,
		},
		{
			Name:            "wire",
			TemplatePath:    "./templates/wire.tmpl",
			FileDestination: utils.Wire,
			ConfigModel:     config,
			Extension:       utils.Go,
		},
		{
			Name:            "wire",
			TemplatePath:    "./templates/wire.tmpl",
			FileDestination: utils.Root,
			ConfigModel:     config,
			Extension:       utils.Go,
		},
		{
			Name:            "Dockerfile",
			TemplatePath:    "./templates/docker/Dockerfile.tmpl",
			FileDestination: utils.Root,
			ConfigModel:     config,
			Extension:       utils.None,
		},
		{
			Name:            "docker-compose",
			TemplatePath:    "./templates/docker/docker-compose.tmpl",
			FileDestination: utils.Root,
			ConfigModel:     config,
			Extension:       utils.Yaml,
		},
		{
			Name:            "Makefile",
			TemplatePath:    "./templates/Makefile.tmpl",
			FileDestination: utils.Root,
			ConfigModel:     config,
			Extension:       utils.None,
		},
		// API Files
		{
			Name:            "server",
			TemplatePath:    "./templates/api/server.tmpl",
			FileDestination: utils.API,
			ConfigModel:     config,
			Extension:       utils.Go,
		},
		{
			Name:            "router",
			TemplatePath:    "./templates/api/router.tmpl",
			FileDestination: utils.API,
			ConfigModel:     config.HTTP,
			Extension:       utils.Go,
		},
		// Config Files
		{
			Name:            "config",
			TemplatePath:    "./templates/config/config.tmpl",
			FileDestination: utils.Config,
			ConfigModel:     config.AppConfig,
			Extension:       utils.Go,
		},
		{
			Name:            "defaults",
			TemplatePath:    "./templates/config/defaults.tmpl",
			FileDestination: utils.Config,
			ConfigModel:     config.AppConfig,
			Extension:       utils.Go,
		},
		{
			Name:            "utils",
			TemplatePath:    "./templates/config/utils.tmpl",
			FileDestination: utils.Config,
			ConfigModel:     config.AppConfig,
			Extension:       utils.Go,
		},
	}

	// Generate packages and groups of files
	handlerFiles := make([]FileDataParams, 0)
	for _, handler := range config.HTTP {
		handlerFiles = append(handlerFiles, FileDataParams{
			Name:            handler.Name,
			TemplatePath:    "./templates/handler.tmpl",
			FileDestination: utils.Handler,
			ConfigModel:     handler,
			HasConstructor:  true,
			Extension:       utils.Go,
		})
	}

	modelFiles := make([]FileDataParams, 0)
	for _, model := range config.Models {
		modelFiles = append(modelFiles, FileDataParams{
			Name:            model.Name,
			TemplatePath:    "./templates/model.tmpl",
			FileDestination: utils.Model,
			ConfigModel:     model,
			HasConstructor:  false,
			Extension:       utils.Go,
		})
	}

	GenerateAll(filesToGenerate)
	GeneratePackage("handlers", handlerFiles, utils.Handler)
	GeneratePackage("models", modelFiles, utils.Model)
	return nil
}
