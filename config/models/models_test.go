package configmodels_test

import (
	"reflect"
	"testing"

	"github.com/colinso/lego/config"
	configmodels "github.com/colinso/lego/config/models"
)

const testYamlPath = "./test.yaml"

func TestTemplateUtils_TypeOf(t *testing.T) {
	expected := configmodels.Base{
		ProjectPath: ".",
		Name:        "testservice",
		AppConfig:   configmodels.App{"AppName": "\"testservice\"", "Host": "\"localhost\"", "Port": "8080", "EnvVar": "\"hellothere\""},
		Models: []configmodels.Model{
			{
				Name: "APIHandlerRequest",
				Fields: map[string]string{
					"id":    "string",
					"value": "int",
				},
			},
		},
	}
	config.ParseConfig(testYamlPath, "./..")
	got := config.GetConfig()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("Expected: %v, got: %v", expected, got)
	}
}
