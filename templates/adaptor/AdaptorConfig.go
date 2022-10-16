package adaptor

var AdaptorConfigTmpl = `package repo

import (
	"fmt"

	"github.com/myste1tainn/hexnet"
	"github.com/spf13/viper"
)

func New{{ .Name }}RepoConfig() *hexnet.Config {
	var key = "{{ .KeyPath }}"
	var cfg hexnet.Config
	if err := viper.UnmarshalKey(key, &cfg); err != nil {
		panic(err)
	}
	fmt.Printf("[info] config with key = %s, loaded cfg = %v\n", key, cfg)
	return &cfg
}
`
