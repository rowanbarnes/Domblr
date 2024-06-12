package widget

import (
	"bytes"
	_ "embed"
)

type Page struct {
	Title string
	Body  Widget
	Style *Style
}

//go:embed script.html
var scriptStr string

//go:embed style.css
var styleStr string

func (page *Page) Setup(style *Style) {
	if page.Style == nil {
		page.Style = style
	}

	page.Body.Setup(page.Style)
}

func (page *Page) Design(buffer *bytes.Buffer) *bytes.Buffer {
	page.Body.Design(buffer)

	return buffer
}

func (page *Page) Render(buffer *bytes.Buffer) *bytes.Buffer {
	buffer.WriteString(`<!DOCTYPE html>`)
	buffer.WriteString(`<html lang="">`)
	buffer.WriteString(`<head>`)
	buffer.WriteString(`<title>`)
	buffer.WriteString(page.Title)
	buffer.WriteString(`</title>`)
	buffer.WriteString("<script src=\"wasm_exec.js\"></script>")
	buffer.WriteString("<style>")
	buffer.WriteString(styleStr)
	page.Design(buffer)
	buffer.WriteString("</style>")
	buffer.WriteString(`</head>`)
	buffer.WriteString(`<body>`)
	page.Body.Render(buffer)
	buffer.WriteString(scriptStr)
	buffer.WriteString(`</body>`)
	buffer.WriteString(`</html>`)

	return buffer
}
