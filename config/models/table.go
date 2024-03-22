package configmodels

type Table struct {
	TableName string `yaml:"tableName"`
	Model     any    `yaml:"model"`
}

func (c Table) Validate() {
	ValidateFields(c)
}
