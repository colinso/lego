package configmodels

type GeneratorConfig struct {
	ProjectPath string
	Name        string              `yaml:"name"`
	AppConfig   AppConfig           `yaml:"appConfig"`
	Models      []ModelConfig       `yaml:"models"`
	HTTP        []HTTPHandlerConfig `yaml:"http"`
	Logic       []LogicConfig       `yaml:"logic"`
}

func (c GeneratorConfig) Validate() {
	ValidateFields(c)
}
