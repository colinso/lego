package configmodels

type ServiceYaml struct {
	Name    string       `yaml:"name"`
	Methods []MethodYaml `yaml:"methods"`
}

type Service struct {
	Name    string
	Methods []Method
}

func (c Service) Validate() {
	ValidateFields(c)
}
