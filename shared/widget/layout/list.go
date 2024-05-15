package layout

import (
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
	Children    []widget.Widget
	Style       *widget.Style
}

func (list *List) Render(buffer *bytes.Buffer) *bytes.Buffer {
	buffer.WriteString(`<div style="display: flex; align-items: stretch; flex-direction: `)
	if list.Axis == ROW {
		buffer.WriteString("row")
	} else if list.Axis == COL {
		buffer.WriteString("column")
	}
	buffer.WriteString("\">")
	for i := range list.Children {
		list.Children[i].Render(buffer)
	}
	buffer.WriteString(`</div>`)
	return buffer
}

func (list *List) Setup(style *widget.Style) {
	widget.Inherit(&list.Style, style)

	for i := 0; i < list.ItemCount; i++ {
		child := list.ItemBuilder(i)
		child.Setup(list.Style)
		list.Children = append(list.Children, child)
	}
}
