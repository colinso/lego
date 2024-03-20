package generator

import (
	"os"

	"github.com/colinso/lego/actions/config"
	"github.com/colinso/lego/actions/utils"
	"github.com/colinso/lego/configmodels"
)

const logicTemplatePath = "./templates/logic.tmpl"

type LogicGenerator struct {
	cfg configmodels.GeneratorConfig
}

func NewLogicGenerator() *LogicGenerator {
	return &LogicGenerator{
		cfg: config.GetConfig(),
	}
}

func (c LogicGenerator) Generate() error {
	logicTemplate, err := os.ReadFile(logicTemplatePath)
	if err != nil {
		return err
	}

	// TODO: Make this package name configurable
	names := make([]string, 0)
	for _, v := range c.cfg.Logic {
		names = append(names, "New"+v.Name)
	}
	NewComponentGenerator().Generate("logic", names, utils.Logic)

	for _, l := range c.cfg.Logic {
		generateFile(l.Name, string(logicTemplate), l, utils.Logic)
	}
	return nil
}
