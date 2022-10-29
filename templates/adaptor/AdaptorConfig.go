package adaptor

var AdaptorConfigTmpl = `package repo

import (
	"fmt"

	"github.com/myste1tainn/msnet"
	"github.com/spf13/viper"
)

func New{{ .Name }}RepoConfig() *msnet.Config {
	var key = "{{ .KeyPath }}"
	var cfg msnet.Config
	if err := viper.UnmarshalKey(key, &cfg); err != nil {
		panic(err)
	}
	fmt.Printf("[info] config with key = %s, loaded cfg = %v\n", key, cfg)
	return &cfg
}
`
