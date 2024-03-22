package generator

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/colinso/lego/utils"
)

const componentTemplateLocation = "./templates/component.tmpl"

// Generates a group of files to a single location with a component.go file
func GeneratePackage(packageName string, data []FileData, ftype utils.FileLocation) {
	constructorNames := make([]string, 0)
	for _, d := range data {
		if d.HasConstructor {
			constructorNames = append(constructorNames, d.Name)
		}
	}

	componentConfig := struct {
		PackageName string
		Objects     []string
	}{
		PackageName: packageName,
		Objects:     constructorNames,
	}

	data = append(data, FileData{
		Name:            "component",
		TemplatePath:    componentTemplateLocation,
		ConfigModel:     componentConfig,
		FileDestination: ftype,
		Extension:       utils.Go,
	})
	GenerateAll(data)
}

// Generates a group of files
func GenerateAll(files []FileData) error {
	errs := make([]string, 0)
	for _, file := range files {
		err := GenerateFile(file)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}
	return errors.New("Errors: " + strings.Join(errs, ","))
}

// Generates a single file
func GenerateFile(file FileData) error {
	templateFile, err := os.ReadFile(file.TemplatePath)
	if err != nil {
		return err
	}

	tmpl, err := template.New(file.Name).Funcs(utils.TemplateFuncs).Parse(string(templateFile))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	f := utils.CreateFileForType(file.FileDestination, file.Name, file.Extension)

	err = tmpl.Execute(f, file.ConfigModel)
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
