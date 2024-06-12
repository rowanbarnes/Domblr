package layout

import (
	"Domblr/shared/comm"
	"Domblr/shared/widget"
	"bytes"
)

// FlexList TODO rename
type FlexList struct {
	ItemCount   int
	ItemBuilder func(int) widget.Widget
	Axis        int
	Children    []widget.Widget
	Style       *widget.Style
	id          int
}

func (list *FlexList) Setup(style *widget.Style) {
	widget.Inherit(&list.Style, style)
	list.id = comm.RegisterElement()

	for i := len(list.Children); i < list.ItemCount; i++ {
		child := list.ItemBuilder(i)
		list.Children = append(list.Children, child)
	}

	for _, child := range list.Children {
		child.Setup(list.Style)
	}
}

func (list *FlexList) Design(buffer *bytes.Buffer) *bytes.Buffer {
	list.Style.Design(buffer, list.id, "",
		map[int]string{
			widget.Background: "background-color",
		}, map[string]string{
			"justify-content": "space-between",
			"align-items":     "center",
		},
	)
	for i := range list.Children {
		list.Children[i].Design(buffer)
	}

	return buffer
}

func (list *FlexList) Render(buffer *bytes.Buffer) *bytes.Buffer {
	widget.OpenTag(buffer, "div", "", list.id, false)
	for i := range list.Children {
		widget.Render(list.Children[i], buffer)
	}
	widget.CloseTag(buffer, "div")

	return buffer
}
