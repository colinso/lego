package configmodels

type Service struct {
	Name    string   `yaml:"name"`
	Methods []Method `yaml:"methods"`
}

func (c Service) Validate() {
	ValidateFields(c)
}
