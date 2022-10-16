package port

var AdaptorTmpl = `package port

import (
    "github.com/gin-gonic/gin"
    "github.com/myste1tainn/hexfnd"
)

type {{ .Name }}Repo interface {
{{- range $fn := .Fns }}
	{{ $fn }}(request {{ $.Name }}{{ $fn }}Request, ctx gin.Context, rctx *hexfnd.RouteContext) (*{{ $.Name }}{{ $fn }}Response, error)
{{- end }}
}
{{ range $fn := .Fns }}
type {{ $.Name }}{{ $fn }}Request struct {
}

type {{ $.Name }}{{ $fn }}Response struct {
}
{{ end }}
`
