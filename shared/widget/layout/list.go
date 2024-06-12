package layout

import (
	"Domblr/shared/comm"
	"Domblr/shared/widget"
	"bytes"
)

const (
	ROW = iota
	COL
)

type List struct {
	ItemCount   int
	ItemBuilder func(int) widget.Widget
	Axis        int
	Style       *widget.Style
	children    []widget.Widget
	id          int
}

func (list *List) Setup(style *widget.Style) {
	widget.Inherit(&list.Style, style)
	list.id = comm.RegisterElement()

	for i := 0; i < list.ItemCount; i++ {
		child := &ListItem{Body: list.ItemBuilder(i)}
		list.children = append(list.children, child)
		child.Setup(list.Style)
	}
}

func (list *List) Design(buffer *bytes.Buffer) *bytes.Buffer {
	list.Style.Design(buffer, list.id, "", map[int]string{}, map[string]string{
		"display":    "flex",
		"list-style": "none",
	})
	for i := range list.children {
		list.children[i].Design(buffer)
	}

	return buffer
}

func (list *List) Render(buffer *bytes.Buffer) *bytes.Buffer {
	widget.OpenTag(buffer, "ul", "", list.id, false)
	for i := range list.children {
		list.children[i].Render(buffer)
	}
	widget.CloseTag(buffer, "ul")

	return buffer
}
