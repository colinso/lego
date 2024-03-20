package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type configSetter struct {
	missingEnvList []string
}

func (c configSetter) SetStringEnv(key string, cfgVar *string) {
	envVar := os.Getenv(key)
	if envVar == "" {
		c.addEnvError(key)
	}
	*cfgVar = envVar
}

func (c configSetter) SetIntEnv(key string, cfgVar *int) {
	envVar := os.Getenv(key)
	if envVar == "" {
		c.addEnvError(key)
	}
	intVar, err := strconv.Atoi(envVar)
	if err != nil {
		c.addEnvError(key)
	}
	*cfgVar = intVar
}

func (c configSetter) addEnvError(key string) {
	c.missingEnvList = append(c.missingEnvList, key)
}

func (c configSetter) getErrors() error {
	if len(c.missingEnvList) > 0 {
		return fmt.Errorf("error retrieving envs: %s", strings.Join(c.missingEnvList, " "))
	}
	return nil
}
