package configmodels

type BaseYaml struct {
	ProjectPath string
	Name        string        `yaml:"name"`
	AppConfig   App           `yaml:"appConfig"`
	Models      []Model       `yaml:"models"`
	HTTP        []HTTPHandler `yaml:"http"`
	Services    []ServiceYaml `yaml:"services"`
	Database    Database      `yaml:"db"`
}

type Base struct {
	ProjectPath string
	Name        string
	AppConfig   App
	Models      []Model
	HTTP        []HTTPHandler
	Services    []Service
	Database    Database
}

func (c Base) Validate() {
	// ValidateFields(c)
	c.AppConfig["DBPassword"] = c.Database.Password
	c.AppConfig["DBUser"] = c.Database.User
	c.AppConfig["DBName"] = c.Database.Name
	c.AppConfig.Validate()
	// c.Database.Validate()
}
