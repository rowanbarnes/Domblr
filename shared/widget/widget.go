package widget

import (
	"bytes"
	"strconv"
)

// Widget interface
type Widget interface {
	// Setup populates the widget tree
	Setup(style *Style)
	// Design writes the CSS contents to buffer
	Design(buffer *bytes.Buffer) *bytes.Buffer
	// Render method returns HTML string
	Render(buffer *bytes.Buffer) *bytes.Buffer
}

func Setup(widget Widget, style *Style) {
	if widget != nil {
		widget.Setup(style)
	}
}

func Design(widget Widget, buffer *bytes.Buffer) {
	if widget != nil {
		widget.Design(buffer)
	}
}

func Render(widget Widget, buffer *bytes.Buffer) {
	if widget != nil {
		widget.Render(buffer)
	}
}

func OpenTag(buffer *bytes.Buffer, tag string, href string, id int, onclick bool) {
	buffer.WriteString("<")
	buffer.WriteString(tag)
	if href != "" {
		buffer.WriteString(" href=\"")
		buffer.WriteString(href)
		buffer.WriteString("\"")
	}
	buffer.WriteString(" id=\"")
	buffer.WriteString(strconv.Itoa(id))
	buffer.WriteString("\" class=\"s")
	buffer.WriteString(strconv.Itoa(id))
	buffer.WriteString("\"")
	if onclick {
		buffer.WriteString("onclick=\"wasm.exports.onclick(")
		buffer.WriteString(strconv.Itoa(id))
		buffer.WriteString(")\"")
	}
	buffer.WriteString(">")
}

func CloseTag(buffer *bytes.Buffer, tag string) {
	buffer.WriteString("</")
	buffer.WriteString(tag)
	buffer.WriteString(">")
}
