package widget

import (
	"bytes"
	"strconv"
)

type Structure struct {
}

func (s *Structure) Render(html *bytes.Buffer) {

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
