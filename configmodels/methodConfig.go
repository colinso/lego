package configmodels

type MethodConfig struct {
	Name    string            `yaml:"name"`
	Accepts map[string]string `yaml:"accepts"`
	Returns []string          `yaml:"returns"`
}

func (c MethodConfig) Validate() {
	ValidateFields(c)
}
