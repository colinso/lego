package configmodels

import (
	"github.com/colinso/lego/utils/datastructures"
	"gopkg.in/yaml.v3"
)

type MethodYaml struct {
	Name    string    `yaml:"name"`
	Accepts yaml.Node `yaml:"accepts"`
	Returns string    `yaml:"returns"`
}

type Method struct {
	Name    string
	Accepts datastructures.OrderedMap[string, string]
	Returns string
}

func (c Method) Validate() {
	ValidateFields(c)
}
