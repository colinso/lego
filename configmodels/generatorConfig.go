package configmodels

type GeneratorConfig struct {
	ProjectPath string
	Name        string              `yaml:"name"`
	AppConfig   AppConfig           `yaml:"appConfig"`
	Models      []ModelConfig       `yaml:"models"`
	HTTP        []HTTPHandlerConfig `yaml:"http"`
	Logic       []LogicConfig       `yaml:"logic"`
	Database    DBConfig            `yaml:"db"`
}

func (c GeneratorConfig) Validate() {
	ValidateFields(c)
	c.AppConfig.Validate()
	c.Database.Validate()
}
