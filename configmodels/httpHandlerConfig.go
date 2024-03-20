package configmodels

type HTTPHandlerConfig struct {
	Name         string `yaml:"handlerName"`
	Method       string `yaml:"method"`
	Path         string `yaml:"path"`
	RequestBody  string `yaml:"requestBody"`
	ResponseBody string `yaml:"responseBody"`
	Function     string `yaml:"function"`
}

func (c HTTPHandlerConfig) Validate() {
	ValidateFields(c)
}
