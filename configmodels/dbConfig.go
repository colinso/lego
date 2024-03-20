package configmodels

type DBConfig struct {
	Type     string
	User     string        `yaml:"user"`
	Password string        `yaml:"password"`
	Name     string        `yaml:"name"`
	Schema   []TableConfig `yaml:"schema"`
}

func (c DBConfig) Validate() {
	ValidateFields(c)
}
