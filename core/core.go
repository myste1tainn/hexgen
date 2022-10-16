package core

import (
	"io/ioutil"
	"os"
	"strings"
)

type AdaptorConfig struct {
	Name           string         `mapstructure:"name"`
	BaseUrl        string         `mapstructure:"baseUrl"`
	ReadTimeout    int            `mapstructure:"readTimeout"`
	ConnectTimeout int            `mapstructure:"connectTimeout"`
	Apis           map[string]Api `mapstructure:"apis"`
}

type Api struct {
	Name   string `mapstructure:"name"`
	Uri    string `mapstructure:"uri"`
	Method string `mapstructure:"method"`
}

type AdaptorTemplateValues struct {
	Module  string
	Name    string
	KeyPath string
	Fns     []string
}

func GetModule() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	gomod, err := ioutil.ReadFile(dir + "/go.mod")
	if err != nil {
		panic(err)
	}
	gomodString := string(gomod)
	moduleLine := strings.Split(gomodString, "\n")[0]
	module := strings.Replace(moduleLine, "module ", "", 1)
	return module
}
