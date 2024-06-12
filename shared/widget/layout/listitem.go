package layout

import (
	"Domblr/shared/widget"
	"bytes"
)

type ListItem struct {
	Body  widget.Widget
	Style *widget.Style
}

func (listitem *ListItem) Setup(style *widget.Style) {
	widget.Inherit(&listitem.Style, style)
	widget.Setup(listitem.Body, listitem.Style)
}

func (listitem *ListItem) Design(buffer *bytes.Buffer) *bytes.Buffer {
	widget.Design(listitem.Body, buffer)
	return buffer
}

func (listitem *ListItem) Render(buffer *bytes.Buffer) *bytes.Buffer {
	widget.OpenTag(buffer, "li", "", -1, false)
	widget.Render(listitem.Body, buffer)
	widget.CloseTag(buffer, "li")

	return buffer
}
