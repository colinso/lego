package generator

import (
	"lego/actions/config"
	"lego/actions/utils"
	"lego/configmodels"
	"os"
)

const handlerTemplatePath = "./templates/handler.tmpl"
const routerTemplatePath = "./templates/api/router.tmpl"

type HandlerGenerator struct {
	cfg configmodels.GeneratorConfig
}

func NewHandlerGenerator() *HandlerGenerator {
	return &HandlerGenerator{
		cfg: config.GetConfig(),
	}
}

func (c HandlerGenerator) Generate() error {
	handlerTemplate, err := os.ReadFile(handlerTemplatePath)
	if err != nil {
		return err
	}

	routerTemplate, err := os.ReadFile(routerTemplatePath)
	if err != nil {
		return err
	}

	NewComponentGenerator().Generate("handlers", getHandlerNames(c.cfg.HTTP), utils.Handler)

	for _, h := range c.cfg.HTTP {
		generateFile(h.Name, string(handlerTemplate), h, utils.Handler)
	}
	generateFile("router", string(routerTemplate), c.cfg.HTTP, utils.API)
	return nil
}

func getHandlerNames(handlers []configmodels.HTTPHandlerConfig) []string {
	handlerNames := make([]string, 0)
	for _, h := range handlers {
		handlerNames = append(handlerNames, h.Name+"Handler")
	}
	return handlerNames
}
