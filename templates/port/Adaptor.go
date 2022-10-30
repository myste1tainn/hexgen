package port

var AdaptorTmpl = `package port

import (
    "github.com/myste1tainn/msfnd"
)

type {{ .Name }}Repo interface {
{{- range $fn := .Fns }}
	{{ $fn }}(request {{ $.Name }}{{ $fn }}Request, rctx *msfnd.RouteContext) (*{{ $.Name }}{{ $fn }}Response, error)
{{- end }}
}
{{ range $fn := .Fns }}
type {{ $.Name }}{{ $fn }}Request struct {
}

type {{ $.Name }}{{ $fn }}Response struct {
}
{{ end }}
`
