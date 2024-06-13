package layout

import (
	"Domblr/shared/widget"
	"bytes"
)

// FlexList TODO rename
type FlexList struct {
	widget.Node
	Items       []widget.Widget
	ItemCount   int
	ItemBuilder func(int) widget.Widget
	Axis        int
}

func (list *FlexList) Setup(parent *widget.Node, id int) {
	// Construct the required children and add them to the node
	list.Node.AddChild(list.Items...)
	for i := len(list.Items); i < list.ItemCount; i++ {
		list.Node.AddChild(list.ItemBuilder(i))
	}

	// Set up the node
	list.Node.Setup(parent, id)
}

func (list *FlexList) Design(css *bytes.Buffer, html *bytes.Buffer) {
	// Create local variables for dependent properties
	flexDir := "row"
	if list.Axis == COL {
		flexDir = "column"
	}

	// Design the list styles
	list.Style.Design(buffer, list.id, "",
		map[int]string{
			widget.Background: "background-color",
		}, map[string]string{
			"justify-content": "space-between",
			"align-items":     "center",
			"flex-direction":  flexDir,
		},
	)

	// Render the node
	list.Node.Render(css, html)
}

func (list *FlexList) Render(buffer *bytes.Buffer) {
	widget.OpenTag(buffer, "div", "", list.id, false)
	for i := range list.Children {
		widget.Render(list.Children[i], buffer)
	}
	widget.CloseTag(buffer, "div")

	return
}
