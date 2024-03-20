package generator

import (
	"fmt"
	"text/template"

	"github.com/colinso/lego/actions/utils"
)

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
