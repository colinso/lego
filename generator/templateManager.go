package generator

import (
	"github.com/colinso/lego/config"
	configmodels "github.com/colinso/lego/config/models"
	"github.com/colinso/lego/utils"
)

type TemplateManager struct {
	cfg configmodels.Base
}

func NewCodeGenerator() *TemplateManager {
	return &TemplateManager{
		cfg: config.GetConfig(),
	}
}

type FileData struct {
	Name            string
	Extension       utils.FileExtension
	TemplatePath    string
	FileDestination utils.FileLocation
	ConfigModel     any
	HasConstructor  bool
}

func (m TemplateManager) Generate() error {
	config := m.cfg
	// Generate individual files
	filesToGenerate := []FileData{
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
		// DB Files
		{
			Name:            "connection",
			TemplatePath:    "./templates/db/connection.tmpl",
			FileDestination: utils.DB,
			ConfigModel:     config.Database,
			Extension:       utils.Go,
		},
	}

	// Generate packages and groups of files
	handlerFiles := make([]FileData, 0)
	for _, handler := range config.HTTP {
		handlerFiles = append(handlerFiles, FileData{
			Name:            handler.Name,
			TemplatePath:    "./templates/handler.tmpl",
			FileDestination: utils.Handler,
			ConfigModel:     handler,
			HasConstructor:  true,
			Extension:       utils.Go,
		})
	}

	serviceFiles := make([]FileData, 0)
	for _, service := range config.Services {
		serviceFiles = append(serviceFiles, FileData{
			Name:            service.Name,
			TemplatePath:    "./templates/service.tmpl",
			FileDestination: utils.Service,
			ConfigModel:     service,
			HasConstructor:  true,
			Extension:       utils.Go,
		})
	}

	repoFiles := make([]FileData, 0)
	for _, table := range config.Database.Schema {
		serviceFiles = append(repoFiles, FileData{
			Name:            table.TableName,
			TemplatePath:    "./templates/repo.tmpl",
			FileDestination: utils.Repo,
			ConfigModel:     table,
			HasConstructor:  true,
			Extension:       utils.Go,
		})
	}

	modelFiles := make([]FileData, 0)
	for _, model := range config.Models {
		modelFiles = append(modelFiles, FileData{
			Name:            model.Name,
			TemplatePath:    "./templates/model.tmpl",
			FileDestination: utils.Model,
			ConfigModel:     model,
			HasConstructor:  false,
			Extension:       utils.Go,
		})
	}

	GenerateAll(filesToGenerate)
	GeneratePackage("handler", handlerFiles, utils.Handler)
	GeneratePackage("service", serviceFiles, utils.Service)
	GeneratePackage("models", modelFiles, utils.Model)
	GeneratePackage("repo", repoFiles, utils.Repo)
	return nil
}
