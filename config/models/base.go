package configmodels

type Base struct {
	ProjectPath string
	Name        string        `yaml:"name"`
	AppConfig   App           `yaml:"appConfig"`
	Models      []Model       `yaml:"models"`
	HTTP        []HTTPHandler `yaml:"http"`
	Logic       []Service     `yaml:"services"`
	Database    Database      `yaml:"db"`
}

func (c Base) Validate() {
	ValidateFields(c)
	c.AppConfig.Validate()
	c.Database.Validate()
}
