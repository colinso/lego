package configmodels

type HTTPHandler struct {
	Name         string `yaml:"handlerName"`
	Method       string `yaml:"method"`
	Path         string `yaml:"path"`
	RequestBody  string `yaml:"requestBody"`
	ResponseBody string `yaml:"responseBody"`
	Service      string `yaml:"service"`
}

func (c HTTPHandler) Validate() {
	ValidateFields(c)
}
