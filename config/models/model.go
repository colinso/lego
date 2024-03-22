package configmodels

type Model struct {
	Name   string            `yaml:"name"`
	Fields map[string]string `yaml:"fields"`
}

func (c Model) Validate() {
	ValidateFields(c)
}
