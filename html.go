package hypermedia

import (
	"io"
	"text/template"
)

const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
	</head>
	<body>
		{{range .Links}}<a rel="{{.Rel}}" href="{{.URL}}">{{ .Name }}</a><br />{{end}}
	</body>
</html>`

var t, _ = template.New("webpage").Parse(tpl)

func RenderHTML(w io.Writer, h Hypermedia) error {
	return t.Execute(w, h)
}
