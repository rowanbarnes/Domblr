package layout

import (
	"Domblr/shared/widget"
	"Domblr/util"
)

// FlexList TODO rename or merge with List somehow
type FlexList struct {
	widget.Node
	Items       []widget.Widget
	ItemCount   int
	ItemBuilder func(int) widget.Widget
	Axis        int
}

func (l *FlexList) Setup(parent *widget.Node, id int) {
	// Set up the node
	l.Node = widget.Node{
		Structure: widget.Structure{
			Tag: "div",
		},
		Style: &widget.Style{
			Properties: map[string]map[string]any{
				"": {
					"background-color": widget.Background,
					"justify-content":  "space-between",
					"align-items":      "center",
					"flex-direction":   util.If(l.Axis == ROW, "row", "column"),
				},
			},
			Constraint: widget.Constraint{
				Width:  util.If(l.Axis == COL, widget.SHRINK, widget.EXPAND),
				Height: util.If(l.Axis == ROW, widget.SHRINK, widget.EXPAND),
			},
		},
	}

	// Construct the required children and add them to the node
	l.Node.AddChild(l.Items...)
	for i := len(l.Items); i < l.ItemCount; i++ {
		l.Node.AddChild(l.ItemBuilder(i))
	}

	// Set up the node
	l.Node.Setup(parent, id)
}
