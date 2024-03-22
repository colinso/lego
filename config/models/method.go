package configmodels

type Method struct {
	Name    string            `yaml:"name"`
	Accepts map[string]string `yaml:"accepts"`
	Returns []string          `yaml:"returns"`
}

func (c Method) Validate() {
	ValidateFields(c)
}
