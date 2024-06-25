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
	Src   string
}

func (s *Structure) Render(ctx *RenderContext, n *Node) {
	if s.Tag != "" {
		s.openTag(ctx.HTML, n.ID)
	}
	ctx.HTML.WriteString(s.Inner)
	for _, child := range n.Children {
		child.Render(ctx)
	}
	if s.Tag != "" {
		s.closeTag(ctx.HTML)
	}

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
	if s.Src != "" {
		html.WriteString(" src=\"")
		html.WriteString(s.Src)
		html.WriteString("\"")
	}
	html.WriteString(">")
}

func (s *Structure) closeTag(html *bytes.Buffer) {
	html.WriteString("</")
	html.WriteString(s.Tag)
	html.WriteString(">")
}
