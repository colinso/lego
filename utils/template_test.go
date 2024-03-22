package utils

import (
	"reflect"
	"testing"

	"github.com/colinso/lego/config"
	configmodels "github.com/colinso/lego/config/models"
	"github.com/colinso/lego/utils/datastructures"
)

func TestTemplateUtils_TypeOf(t *testing.T) {
	type test struct {
		description string
		input       string
		want        string
	}

	tests := []test{
		{description: "StringSuccess", input: "\"a/b/c\"", want: "string"},
		{description: "BoolSuccessTrue", input: "true", want: "bool"},
		{description: "BoolSuccessFalse", input: "false", want: "bool"},
		{description: "IntSuccess", input: "1234", want: "int"},
		{description: "FloatSuccess", input: "12.34", want: "float64"},
		{description: "ObjectSuccess", input: "TestObject", want: "TestObject"},
	}

	for _, tc := range tests {
		got := TypeOf(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Test %v expected: %v, got: %v", tc.description, tc.want, got)
		}
	}
}

func TestTemplateUtils_ZeroValue(t *testing.T) {
	type test struct {
		description string
		input       string
		want        string
	}

	tests := []test{
		{description: "StringSuccess", input: "string", want: "\"\""},
		{description: "BoolSuccessTrue", input: "bool", want: "false"},
		{description: "IntSuccess", input: "int", want: "0"},
		{description: "FloatSuccess", input: "float64", want: "0"},
		{description: "ObjectSuccess", input: "TestObject", want: "TestObject{}"},
		{description: "PointerSuccess", input: "*TestObject", want: "nil"},
	}

	for _, tc := range tests {
		got := ZeroValue(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Test %v expected: %v, got: %v", tc.description, tc.want, got)
		}
	}
}

func TestTemplateUtils_GetMethodSignatureByValue(t *testing.T) {
	type test struct {
		input configmodels.Method
		want  string
	}

	accepts := datastructures.NewOrderedMap[string, string]()
	accepts.Set("arg1", "string")
	accepts.Set("arg2", "bool")
	accepts.Set("arg3", "TestConfigObject")

	tests := []test{
		{input: configmodels.Method{
			Name:    "TestMethod",
			Accepts: *accepts,
			Returns: "int",
		}, want: "TestMethod(arg1 string, arg2 bool, arg3 TestConfigObject) (int, error)"},
	}

	for _, tc := range tests {
		got := GetMethodSignatureByValue(tc.input.Name, tc.input.Accepts, tc.input.Returns)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestTemplateUtils_GetMethodSignatureByName(t *testing.T) {
	accepts1 := datastructures.NewOrderedMap[string, string]()
	accepts1.Set("arg", "string")
	accepts1.Set("something", "float64")

	accepts2 := datastructures.NewOrderedMap[string, string]()
	accepts2.Set("arg", "bool")
	accepts2.Set("something", "float64")

	config.SetConfig(configmodels.Base{
		Services: []configmodels.Service{
			{
				Name: "TestClass",
				Methods: []configmodels.Method{
					{
						Name:    "Add",
						Accepts: *accepts1,
						Returns: "int",
					},
					{
						Name:    "Subtract",
						Accepts: *accepts2,
						Returns: "",
					},
				},
			},
		},
	})

	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "TestClass.Add", want: "Add(arg string, something float64) (int, error)"},
		{input: "TestClass.Subtract", want: "Subtract(arg bool, something float64) error"},
	}

	for _, tc := range tests {
		got := GetMethodSignatureByName(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestTemplateUtils_GetHandlerLogicMethodString(t *testing.T) {
	accepts1 := datastructures.NewOrderedMap[string, string]()
	accepts1.Set("arg", "string")
	accepts1.Set("something", "float64")

	accepts2 := datastructures.NewOrderedMap[string, string]()
	accepts2.Set("arg", "bool")
	accepts2.Set("something", "float64")

	config.SetConfig(configmodels.Base{
		Services: []configmodels.Service{
			{
				Name: "TestClass",
				Methods: []configmodels.Method{
					{
						Name:    "Add",
						Accepts: *accepts1,
						Returns: "int",
					},
					{
						Name:    "Subtract",
						Accepts: *accepts2,
						Returns: "",
					},
				},
			},
		},
	})

	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "TestClass.Add", want: "response, err := h.cmd.Add(body.Arg, body.Something)"},
		{input: "TestClass.Subtract", want: "response, err := h.cmd.Subtract(body.Arg, body.Something)"},
	}

	for _, tc := range tests {
		got := GetHandlerLogicMethodString(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
