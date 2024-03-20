package configmodels

type ModelConfig struct {
	Name   string            `yaml:"name"`
	Fields map[string]string `yaml:"fields"`
}

func (c ModelConfig) Validate() {
	ValidateFields(c)
}
