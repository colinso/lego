package configmodels

type Table struct {
	TableName string   `yaml:"tableName"`
	Model     string   `yaml:"model"`
	Ops       []string `yaml:"ops"`
}

func (c Table) Validate() {
	ValidateFields(c)
}
