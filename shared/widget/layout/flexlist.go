package layout

import (
	"Domblr/shared/widget"
	"Domblr/util"
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
	// Set up the node
	list.Node = widget.Node{
		Structure: widget.Structure{
			Tag: "div",
		},
		Style: &widget.Style{
			Transform: map[int]string{
				widget.Background: "background-color",
			},
			Custom: map[string]map[string]string{
				"": {
					"justify-content": "space-between",
					"align-items":     "center",
					"flex-direction":  util.If(list.Axis == ROW, "row", "column"),
				},
			},
			Constraint: widget.Constraint{
				Width:  util.If(list.Axis == COL, widget.SHRINK, widget.EXPAND),
				Height: util.If(list.Axis == ROW, widget.SHRINK, widget.EXPAND),
			},
		},
	}

	// Construct the required children and add them to the node
	list.Node.AddChild(list.Items...)
	for i := len(list.Items); i < list.ItemCount; i++ {
		list.Node.AddChild(list.ItemBuilder(i))
	}

	// Set up the node
	list.Node.Setup(parent, id)
}
