package main

import (
	"html/template"
)

var indexTemplate = template.Must(template.New("index").Parse(indexTemplateSrc))

const indexTemplateSrc = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <title>Coffee!!</title>
    <meta name="viewport" content="width=device-width; initial-scale=1.0; maximum-scale=1.0; user-scalable=0;" />
  </head>
  <body>
    <h1>Coffee!!</h1>

    {{ if len . }}
    <p>
      {{ range . }}
      {{ . }}
      {{ end }}
    </p>
    {{ end }}

    <p>
      <form method="post" action="/">
        <input type="hidden" name="message" value="Coffee's brewing...">
        <input type="submit" value="Coffee's brewing">
      </form>
    </p>
  </body>
</html>
`
