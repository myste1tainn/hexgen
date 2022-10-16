package adaptor

var AdaptorTmpl = `package repo

import (
    "strings"

    "github.com/gin-gonic/gin"
	"{{ .Module }}/internal/core/port"
    "github.com/myste1tainn/hexfnd"
	"github.com/myste1tainn/hexnet"
)

type {{ .Name }}Repo struct {
	client hexnet.Client
	config *hexnet.Config
}

func New{{ .Name }}Repo(client hexnet.Client, config *hexnet.Config) port.{{ .Name }}Repo {
	return {{ .Name }}Repo{
		client: client,
		config: config,
	}
}
{{ range $fn := .Fns }}
func (r {{ $.Name }}Repo) {{ $fn }}(req port.{{ $.Name }}{{ $fn }}Request, ctx gin.Context, rctx *hexfnd.RouteContext) (*port.{{ $.Name }}{{ $fn }}Response, error) {
    var l = ctx.GetLogger().NewSpanLogger()
	defer l.Destroy()

    var configKey = strings.ToLower("{{ $fn }}")
    var result port.{{ $.Name }}{{ $fn }}Response
	var error hexnet.ErrorResponse
	res, err := r.client.
		WithLogger(l).
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
