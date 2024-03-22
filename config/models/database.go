package configmodels

type Database struct {
	Type     string
	User     string  `yaml:"user"`
	Password string  `yaml:"password"`
	Name     string  `yaml:"name"`
	Schema   []Table `yaml:"schema"`
}

func (c Database) Validate() {
	ValidateFields(c)
}
