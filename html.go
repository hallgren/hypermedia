package hypermedia

import (
	"io"
	"text/template"
)

const tpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <style type="text/css">
      * {
        box-sizing: border-box;
      }

      body {
        background-color: #F9E9DE;
        margin: 0;
        font-family: sans-serif;
      }

      .resource {
        margin-left: 0.3em;
        padding-left: 0.3em;
      }

      body .resource {
        padding-right: 0.5em;
      }

      .resource .resource {
        border-left: 4px solid #bababa;
      }

      h1,h2,h3,h4,h5 {
        background-color: #FFB662;
        padding: 0.2em;
        margin-bottom: 0.2em;
      }

      .attributes {
        margin-bottom: 0.2em;
      }

      th { background-color: #FFD19D; }
      td { background-color: #E2ECBA; }

      .rel a {
        color: black;
        text-decoration: none;
      }

      .subresource {
        padding-top: 1em;
      }

      .subresource-rel {
        background-color: #FFB662;
        padding: 0.5em;
        margin-bottom: 0.5em;
      }

      .header, .footer {
        width: 100%;
        background-color: black;
        color: white;
        padding: 0.3em;
      }

      .footer a {
        text-decoration: none;
        color: white;
      }

      .footer {
        margin-top: 1em;
      }

      .collapsed .heading::before, .expanded .heading::before {
        font-size: 0.6em;
        display: block;
        float: left;
        margin-top: 0.4em;
        margin-right: 0.5em;
      }

      .collapsed .heading::before {
        content: "▶";
      }

      .expanded .heading::before {
        content: "▼";
      }

      .collapsed .body {
        display: none;
      }
    </style>
  </head>
  <body>

	{{ define "form" }}
		<form rel="{{ .Rel }}" action="{{ .URL }}" method="{{ .Method }}">
			{{ range .Inputs }}
				{{ if .Label }}
					<label for="{{ .ID }}">{{ .Label }}:</label><br>
				{{ end }}
				<input type="{{ .Type }}" id="{{ .ID }}" name="{{ .Name }}" value="{{ .Value }}"><br>
			{{ end }}
		</form>
	{{ end }}

	{{ define "r" }}
		<table class="attributes">
        {{range $key, $value := .Properties}}
        	<tr class="attribute"><td class="name">{{ $key }}</td><td class="value">{{ $value }}</td></tr>
        {{end}}
        </table>
		<table class="links">
			<tr>
            	<th>Rel</th>
            	<th>URI</th>
            	<th>Title</th>
            	<th>Templated</th>
            </tr>
            {{range .Links}}
            <tr class="link">
            	<td class="rel"><a href="">{{ .Rel }}</a></td>
                <td class="uri"><a href="{{ .URL }}">{{ .URL }}</a></td>
                <td class="title">{{ .Name }}</td>
                <td class="templated">false</td>
            </tr>
            {{end}}
        </table>

		{{ if gt (len .Forms) 0}}
        	<div class="forms">
			{{ range .Forms }}
				{{ template "form" . }}
			{{ end }}
			</div>
		{{ end }}
	{{ end }}

    <div class="header request-info">
    </div>

    <div class="resource">
      <h1 class="heading type">Resource Type</h1>

      <div class="resource-data body">
		{{ template "r" . }}

		
		{{ if gt (len .Resources) 0}}
			{{ range .Resources }}
        		<div class="subresource collapsed">
          			<div class="subresource-rel rel heading"><a>{{ .Name }}</a></div>
          			<div class="value body" class="collapsed">
						{{ template "r" . }}
		  			</div>
		  		</div>
			{{ end }}
		{{ end }}
      </div>
    </div>

    <script>
      function setupCollapseExpand() {
        var nodes = document.getElementsByClassName("heading");

        for (var i=0; i < nodes.length; i++) {
          var node = nodes[i];
          node.onclick = function() {
            var parent = this.parentNode;
            if (parent.className.indexOf("collapsed") >= 0) {
              parent.className = parent.className.replace("collapsed", "expanded");
            } else {
              parent.className = parent.className.replace("expanded", "collapsed");
            }
          };
        }
      };

      if (window.addEventListener) {
        window.addEventListener("load", setupCollapseExpand, false);
      } else if (window.attachEvent) {
        window.attachEvent("onload", setupCollapseExpand);
      }
    </script>
  </body>
</html>`

var t, _ = template.New("webpage").Parse(tpl)

func RenderHTML(w io.Writer, h *Resource) error {
	return t.Execute(w, h)
}
