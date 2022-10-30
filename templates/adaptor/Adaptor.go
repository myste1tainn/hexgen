package adaptor

var AdaptorTmpl = `package repo

import (
    "strings"

	"{{ .Module }}/internal/core/port"
    "github.com/myste1tainn/msfnd"
	"github.com/myste1tainn/msnet"
)

type {{ .Name }}Repo struct {
	client msnet.Client
	config *msnet.Config
}

func New{{ .Name }}Repo(client msnet.Client, config *msnet.Config) port.{{ .Name }}Repo {
	return {{ .Name }}Repo{
		client: client,
		config: config,
	}
}
{{ range $fn := .Fns }}
func (r {{ $.Name }}Repo) {{ $fn }}(req port.{{ $.Name }}{{ $fn }}Request, rctx *msfnd.RouteContext) (*port.{{ $.Name }}{{ $fn }}Response, error) {
    var configKey = strings.ToLower("{{ $fn }}")
    var result port.{{ $.Name }}{{ $fn }}Response
	var error msnet.ErrorResponse
	res, err := r.client.
		RequestWithContext(rctx, r.config, configKey).
		Call(&result, &error)
	if err != nil {
		return nil, err
	} else if res.IsError() {
        return nil, error
	} else {
		return &result, nil
	}
}
{{ end }}
`
