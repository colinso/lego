package configmodels

import (
	"fmt"
	"strconv"
)

type App map[string]string

func (a App) Validate() {
	_, appNameOk := a["AppName"]
	_, hostOk := a["Host"]
	_, portOk := a["Port"]

	if !appNameOk || !hostOk || !portOk {
		panic("Application configuration must contain 'AppName', 'Host' and 'Port'.")
	}

	for k, v := range a {
		a[k] = convertToTypedString(v)
	}
}

func convertToTypedString(str string) string {
	if isFloat(str) || isInt(str) || isBool(str) {
		return str
	}
	return fmt.Sprintf("\"%v\"", str)
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func isBool(s string) bool {
	return s == "true" || s == "false"
}
