package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"github.com/colinso/lego/config"
	configmodels "github.com/colinso/lego/config/models"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	String  = "string"
	Integer = "int"
	Float   = "float64"
	Boolean = "bool"
)

var TemplateFuncs = template.FuncMap(map[string]any{
	"GetServiceName": func() string { return config.GetConfig().AppConfig["AppName"] },
	"GetModuleName":  func() string { return config.GetConfig().Name },
	"ToConfigCase": func(str string) string {
		var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
		var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
		snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
		snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
		return strings.ToUpper(snake)
	},
	"TypeOf":                      TypeOf,
	"TitleCasedTypeOf":            func(str string) string { return cases.Title(language.English).String(TypeOf(str)) },
	"IsString":                    func(str string) bool { return TypeOf(str) == "string" },
	"FirstToLower":                FirstToLower,
	"GetFirstLetter":              func(str string) string { return string(str[0]) },
	"IsLastIndexInMap":            IsLastIndexInMap[string, string],
	"IsLastIndexInSlice":          IsLastIndexInSlice[string],
	"ZeroValue":                   ZeroValue,
	"GetMethodSignatureByName":    GetMethodSignatureByName,
	"GetMethodSignatureByValue":   GetMethodSignatureByValue,
	"GetFunctionClass":            func(str string) string { return strings.Split(str, ".")[0] },
	"GetHandlerLogicMethodString": GetHandlerLogicMethodString,
})

func GetHandlerLogicMethodString(name string) string {
	m := GetLogicMethod(name)

	argsString := ""
	i := 0
	for k, _ := range m.Accepts {
		argsString += "body." + cases.Title(language.English).String(k)
		if i < len(m.Accepts)-1 {
			argsString += ", "
		}
		i++
	}
	return fmt.Sprintf("response, err := h.cmd.%v(%v)", m.Name, argsString)
}

func GetLogicMethod(name string) configmodels.Method {
	names := strings.Split(name, ".")
	logics := config.GetConfig().Logic
	for _, v := range logics {
		if v.Name == names[0] {
			for _, m := range v.Methods {
				if m.Name == names[1] {
					return m
				}
			}
			break
		}
	}
	panic("Could not find method signature by name: " + name)
}

func GetMethodSignatureByName(name string) string {
	m := GetLogicMethod(name)
	return GetMethodSignatureByValue(m.Name, m.Accepts, m.Returns)
}

func GetMethodSignatureByValue(name string, args map[string]string, returns []string) string {
	i := 0
	argsString := ""
	for k, v := range args {
		argsString += fmt.Sprintf("%v %v", k, v)
		if i < len(args)-1 {
			argsString += ", "
		}
		i++
	}

	returnsString := ""
	if len(returns) > 1 {
		returnsString += "("
	}
	returnsString += strings.Join(returns, ", ")
	if len(returns) > 1 {
		returnsString += ")"
	}

	return fmt.Sprintf("%v(%v) %v", name, argsString, returnsString)
}

func IsLastIndexInMap[K string, V any](m map[K]V, index int) bool {
	return index == len(m)-1
}

func IsLastIndexInSlice[T any](arr []T, index int) bool {
	return index == len(arr)-1
}

func TypeOf(str string) string {
	if strings.Contains(str, "\"") {
		return String
	} else if strings.EqualFold(str, "true") || strings.EqualFold(str, "false") {
		return Boolean
	} else if _, err := strconv.ParseInt(str, 10, 64); err == nil {
		return Integer
	} else if _, err := strconv.ParseFloat(str, 64); err == nil {
		return Float
	} else {
		return str
	}
}

func ZeroValue(valueType string) string {
	if valueType == String {
		return "\"\""
	} else if valueType == Boolean {
		return "false"
	} else if valueType == Integer || valueType == Float {
		return "0"
	} else {
		return valueType + "{}"
	}
}

func FirstToLower(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size <= 1 {
		return s
	}
	lc := unicode.ToLower(r)
	if r == lc {
		return s
	}
	return string(lc) + s[size:]
}
