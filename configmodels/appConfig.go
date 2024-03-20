package configmodels

type AppConfig map[string]string

func (a AppConfig) Validate() {
	_, appNameOk := a["AppName"]
	_, hostOk := a["Host"]
	_, portOk := a["Port"]

	if !appNameOk || !hostOk || !portOk {
		panic("Application configuration must contain 'AppName', 'Host' and 'Port'.")
	}
}
