package configmodels

type TableConfig struct {
	TableName string `yaml:"tableName"`
	Model     any    `yaml:"model"`
}

func (c TableConfig) Validate() {
	ValidateFields(c)
}
