package widget

import (
	"Domblr/util"
	"bytes"
	"strconv"
)

type Structure struct {
	Tag     string
	Href    string
	Onclick bool
	// Inner TODO consider making inner a Node instead of part of Structure
	Inner string
}

func (s *Structure) Render(css *bytes.Buffer, html *bytes.Buffer, n *Node) {
	s.openTag(html, n.id)
	html.WriteString(s.Inner)
	for _, child := range n.Children {
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
	html.WriteString("\" class=\"")
	html.WriteString(util.ItoABase26(id))
	html.WriteString("\"")
	if s.Onclick {
		html.WriteString("onclick=\"wasm.exports.invoke(")
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
