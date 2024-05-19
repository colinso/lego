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
	"github.com/colinso/lego/utils/datastructures"

	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	String  = "string"
	Error   = "error"
	Integer = "int"
	Float   = "float64"
	Boolean = "bool"
)

var TemplateFuncs = template.FuncMap(map[string]any{
	// global configs
	"GetServiceName": func() string { return config.GetConfig().AppConfig["AppName"] },
	"GetModuleName":  func() string { return config.GetConfig().Name },
	// string
	"ToConfigCase": func(str string) string {
		var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
		var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
		snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
		snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
		return strings.ToUpper(snake)
	},
	"TypeOf":           TypeOf,
	"TitleCasedTypeOf": func(str string) string { return cases.Title(language.English).String(TypeOf(str)) },
	"IsString":         func(str string) bool { return TypeOf(str) == "string" },
	"FirstToLower":     FirstToLower,
	"ToTitle":          cases.Title(language.English).String,
	"GetFirstLetter":   func(str string) string { return strings.ToLower(string(str[0])) },
	// arrays
	"IsLastIndexInMap":   IsLastIndexInMap[string, string],
	"IsLastIndexInSlice": IsLastIndexInSlice[string],
	"SliceContains":      slices.Contains[[]string, string],
	"IsNotEmpty":         func(array []configmodels.Service) bool { return len(array) > 0 },
	// values and methods
	"ZeroValue":                   ZeroValue,
	"ZeroReturnValues":            ZeroReturnValues,
	"GetMethodSignatureByName":    GetMethodSignatureByName,
	"GetMethodSignatureByValue":   GetMethodSignatureByValue,
	"GetHandlerLogicMethodString": GetHandlerLogicMethodString,
	"GetServiceClass":             func(str string) string { return strings.Split(str, ".")[0] },
	// repo
	"BuilderFunc": CreateBuilderFunc,
})

func GetHandlerLogicMethodString(name string) string {
	m := GetLogicMethod(name)

	argsString := ""

	for i, k := range m.Accepts.Keys() {
		argsString += "body." + cases.Title(language.English).String(k)
		if i < len(m.Accepts.Keys())-1 {
			argsString += ", "
		}
		i++
	}
	return fmt.Sprintf("response, err := h.service.%v(%v)", m.Name, argsString)
}

func GetLogicMethod(name string) configmodels.Method {
	names := strings.Split(name, ".")
	services := config.GetConfig().Services

	for _, v := range services {
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

func GetMethodSignatureByValue(name string, args datastructures.OrderedMap[string, string], returns string) string {
	i := 0
	argsString := ""
	for _, k := range args.Keys() {
		v, _ := args.Get(k)
		argsString += fmt.Sprintf("%v %v", k, v)
		if i < len(args.Keys())-1 {
			argsString += ", "
		}
		i++
	}

	returnString := ""
	if returns == "" {
		returnString = "error"
	} else {
		returnString = fmt.Sprintf("(%v, error)", returns)
	}

	return fmt.Sprintf("%v(%v) %v", name, argsString, returnString)
}

func ZeroReturnValues(returns string) string {
	if returns == "" {
		return "nil"
	} else {
		return fmt.Sprintf("%v, nil", ZeroValue(returns))
	}
}

func IsLastIndexInMap[K string, V any](m map[K]V, index int) bool {
	return index == len(m)-1
}

func IsLastIndexInSlice[T any](arr []T, index int) bool {
	return index == len(arr)-1
}

func TypeOf(str string) string {
	pointerPrefix := ""
	if strings.HasPrefix(str, "*") {
		pointerPrefix = "*"
	}

	if strings.EqualFold(str, "\"error\"") {
		return pointerPrefix + Error
	} else if strings.Contains(str, "\"") {
		return pointerPrefix + String
	} else if strings.EqualFold(str, "true") || strings.EqualFold(str, "false") {
		return pointerPrefix + Boolean
	} else if _, err := strconv.ParseInt(str, 10, 64); err == nil {
		return pointerPrefix + Integer
	} else if _, err := strconv.ParseFloat(str, 64); err == nil {
		return pointerPrefix + Float
	} else {
		return pointerPrefix + str
	}
}

func ZeroValue(valueType string) string {
	if valueType == String {
		return "\"\""
	} else if valueType == Boolean {
		return "false"
	} else if valueType == Integer || valueType == Float {
		return "0"
	} else if valueType == Error {
		return "nil"
	} else if valueType != "" {
		return valueType + "{}"
	}
	return ""
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
func GetModelByName(modelName string) configmodels.Model {
	modelIndex := slices.IndexFunc(config.GetConfig().Models, func(a configmodels.Model) bool {
		return a.Name == modelName
	})
	return config.GetConfig().Models[modelIndex]
}

func CreateBuilderFunc(table configmodels.Table) string {
	builderFunc := `sqlString, args := builder.
		InsertInto("%s").
		Cols(%s).
		Values(%s).
		BuildWithFlavor(sqlbuilder.PostgreSQL)`
	model := GetModelByName(table.Model)
	cols := ""
	vals := ""
	for k, _ := range model.Fields {
		cols += fmt.Sprintf("\"%s\",", strings.ToLower(k))
		vals += fmt.Sprintf("m.%s,", k)
	}
	cols = strings.TrimRight(cols, ",")
	vals = strings.TrimRight(vals, ",")
	return fmt.Sprintf(builderFunc, table.TableName, cols, vals)
}

func GetBuilderFunc(table configmodels.Table) string {
	builderFunc := `sqlString, args := builder.
		InsertInto("%s").
		Cols(%s).
		Values(%s).
		BuildWithFlavor(sqlbuilder.PostgreSQL)`
	model := GetModelByName(table.Model)
	cols := ""
	vals := ""
	for k, _ := range model.Fields {
		cols += fmt.Sprintf("\"%s\",", strings.ToLower(k))
		vals += fmt.Sprintf("m.%s,", k)
	}
	cols = strings.TrimRight(cols, ",")
	vals = strings.TrimRight(vals, ",")
	return fmt.Sprintf(builderFunc, table.TableName, cols, vals)
}
