package widget

import (
	"bytes"
	"strconv"
)

type Structure struct {
	Tag     string
	Href    string
	Onclick bool
}

func (s *Structure) Render(css *bytes.Buffer, html *bytes.Buffer, n *Node) {
	s.openTag(html, n.id)
	for _, child := range n.children {
		child.Render(css, html)
	}
	s.closeTag(html)
}

func (s *Structure) openTag(html *bytes.Buffer, id int) {
	html.WriteString("<")
	html.WriteString(s.Tag)
	if s.Href != "" {
		html.WriteString(" href=\"")
		html.WriteString(s.Href)
		html.WriteString("\"")
	}
	html.WriteString(" id=\"")
	html.WriteString(strconv.Itoa(id))
	html.WriteString("\" class=\"s")
	html.WriteString(strconv.Itoa(id))
	html.WriteString("\"")
	if s.Onclick {
		html.WriteString("onclick=\"wasm.exports.onclick(")
		html.WriteString(strconv.Itoa(id))
		html.WriteString(")\"")
	}
	html.WriteString(">")
}

func (s *Structure) closeTag(html *bytes.Buffer) {
	html.WriteString("</")
	html.WriteString(s.Tag)
	html.WriteString(">")
}
