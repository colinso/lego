package configmodels

type LogicConfig struct {
	Name    string         `yaml:"name"`
	Methods []MethodConfig `yaml:"methods"`
}

func (c LogicConfig) Validate() {
	ValidateFields(c)
}
