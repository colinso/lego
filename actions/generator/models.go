package generator

import (
	"os"

	"github.com/colinso/lego/actions/config"
	"github.com/colinso/lego/actions/utils"
	"github.com/colinso/lego/configmodels"
)

const modelTemplatePath = "./templates/model.tmpl"

type ModelGenerator struct {
	cfg configmodels.GeneratorConfig
}

func NewModelGenerator() *ModelGenerator {
	return &ModelGenerator{
		cfg: config.GetConfig(),
	}
}

func (c ModelGenerator) Generate() error {
	modelTemplate, err := os.ReadFile(modelTemplatePath)
	if err != nil {
		return err
	}

	for _, m := range c.cfg.Models {
		generateFile(m.Name, string(modelTemplate), m, utils.Model)
	}
	return nil
}
