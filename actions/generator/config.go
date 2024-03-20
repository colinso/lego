package generator

import (
	"fmt"
	"os"
	"strconv"

	"github.com/colinso/lego/actions/config"
	"github.com/colinso/lego/actions/utils"
	"github.com/colinso/lego/configmodels"
)

const configTemplatePath = "./templates/config/config.tmpl"
const defaultsTemplatePath = "./templates/config/defaults.tmpl"
const utilsTemplatePath = "./templates/config/utils.tmpl"

var configTemplatePaths = map[string]string{
	"config":   "./templates/config/config.tmpl",
	"defaults": "./templates/config/defaults.tmpl",
	"utils":    "./templates/config/utils.tmpl",
}

type ConfigGenerator struct {
	cfg configmodels.GeneratorConfig
}

func NewConfigGenerator() *ConfigGenerator {
	return &ConfigGenerator{
		cfg: config.GetConfig(),
	}
}

func (c ConfigGenerator) Generate() error {
	for k, v := range c.cfg.AppConfig {
		c.cfg.AppConfig[k] = convertToTypedString(v)
	}

	for k, v := range configTemplatePaths {
		tmpl, err := os.ReadFile(v)
		if err != nil {
			return err
		}

		generateFile(k, string(tmpl), c.cfg.AppConfig, utils.Config)
	}
	return nil
}

func convertToTypedString(str string) string {
	if isFloat(str) || isInt(str) || isBool(str) {
		return str
	}
	return fmt.Sprintf("\"%v\"", str)
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func isBool(s string) bool {
	return s == "true" || s == "false"
}
