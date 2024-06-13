package layout

import (
	"Domblr/shared/comm"
	"Domblr/shared/widget"
	"bytes"
)

type ListItem struct {
	Body  widget.Widget
	Style *widget.Style
	id    int
}

func (listitem *ListItem) Setup(style *widget.Style) {
	widget.Inherit(&listitem.Style, style)
	listitem.id = comm.RegisterElement()
	widget.Setup(listitem.Body, listitem.Style)
}

func (listitem *ListItem) Design(buffer *bytes.Buffer) widget.Constraint {
	listitem.Style.Design(buffer, listitem.id, "",
		map[int]string{},
		map[string]string{
			"width":  "fit-content",
			"height": "fit-content",
		},
	)
	widget.Design(listitem.Body, buffer)
}

func (listitem *ListItem) Render(buffer *bytes.Buffer) {
	widget.OpenTag(buffer, "li", "", listitem.id, false)
	widget.Render(listitem.Body, buffer)
	widget.CloseTag(buffer, "li")

	return
}
