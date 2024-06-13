package widget

import (
	"Domblr/res"
	"bytes"
	_ "embed"
)

type Page struct {
	Title string
	Body  Widget
	Style *Style
}

func (page *Page) Setup(style *Style) {
	if page.Style == nil {
		page.Style = style
	}

	page.Body.Setup(nil, 0)
}

func (page *Page) Render(buffer *bytes.Buffer) {
	// Render the css and html code
	var css, html bytes.Buffer
	page.Body.Render(&css, &html)

	// Write the boilerplate document and the rendered css/html
	// TODO cleanup boilerplate code, use a template or file of some kind
	buffer.WriteString(`<!DOCTYPE html>`)
	buffer.WriteString(`<html lang="">`)
	buffer.WriteString(`<head>`)
	buffer.WriteString(`<title>`)
	buffer.WriteString(page.Title)
	buffer.WriteString(`</title>`)
	buffer.WriteString(`<script src="wasm_exec.js"></script>`)
	buffer.WriteString(`<style>`)
	buffer.WriteString(res.GlobalStyles)
	buffer.Write(css.Bytes())
	buffer.WriteString(`</style>`)
	buffer.WriteString(`</head>`)
	buffer.WriteString(`<body>`)
	buffer.Write(html.Bytes())
	buffer.WriteString(`<script>`)
	buffer.WriteString(res.LauncherScript)
	buffer.WriteString(`</script>`)
	buffer.WriteString(`</body>`)
	buffer.WriteString(`</html>`)
}
