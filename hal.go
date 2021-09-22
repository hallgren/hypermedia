// This is based on the extension on HAL https://gist.github.com/mikekelly/893552
package hypermedia

import (
	"io"
	"text/template"
)

const haloTPL = `
	{{ define "forms" }}
	{{ end }}

	{{ define "links" }}
		{{ if gt (len .Links) 0 }}
		"_links": {
			{{ $first := true }}
			{{ range .Links }}
				{{ if $first }}
					{{ $first = false }}
				{{ else }}
					,
				{{ end }}
				"{{.REL}}": {
					"href": "{{ .URL }}"
				}
			{{ end }}
		},
		{{ end }}
	{{ end }}

	{{ define "properties" }}
		{{ $first := true }}
        {{range $key, $value := .Properties}}
			{{ if $first }}
				{{ $first = false }}
			{{ else }}
				,
			{{ end }}
			"{{ $key }}":"{{ $value }}"
		{{ end }}
	{{ end }}

	{{ define "resource" }}
		{{ template "links" . }}
		{{ template "forms" . }}
		{{ if gt (len .Resources) 0}}
			{{ range .Resources }}
				"_embedded": {
					"{{ .Name }}": {
						{{ template "resource" . }}
					}
				},
			{{ end }}
		{{ end }}
		{{ template "properties" . }}
	{{ end }}

	{
		{{ template "resource" . }}
	
	}
`

var haloTemplate, _ = template.New("halo").Parse(haloTPL)

func RenderHal(w io.Writer, h *Resource) error {
	return haloTemplate.Execute(w, h)
}
