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

//go:embed head.html
var headStr string

func (page *Page) Setup(style *Style) {
	if page.Style == nil {
		page.Style = style
	}
	page.Body.Setup(page.Style)
}

func (page *Page) Render(buffer *bytes.Buffer) *bytes.Buffer {
	buffer.WriteString(`<!DOCTYPE html>`)
	buffer.WriteString(`<html lang="">`)
	buffer.WriteString(`<head>`)
	buffer.WriteString(`<title>`)
	buffer.WriteString(page.Title)
	buffer.WriteString(`</title>`)
	buffer.WriteString("<script src=\"wasm_exec.js\"></script>")
	buffer.WriteString(`</head>`)
	buffer.WriteString(`<body>`)
	page.Body.Render(buffer)
	buffer.WriteString(`</body>`)
	buffer.WriteString(`</html>`)
	buffer.WriteString(headStr)

	return buffer
}
