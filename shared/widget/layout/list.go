package layout

import (
	"Domblr/shared/widget"
	"Domblr/util"
)

const (
	ROW = iota
	COL
)

type List struct {
	widget.Node
	ItemCount   int
	ItemBuilder func(int) widget.Widget
	Axis        int
}

func (l *List) Setup(parent *widget.Node, id int) {
	// Set up the node
	l.Node = widget.Node{
		Structure: widget.Structure{
			Tag: "ul",
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

	// Construct the children and add them to the node
	for i := 0; i < l.ItemCount; i++ {
		l.Node.AddChild(&ListItem{Body: l.ItemBuilder(i)})
	}

	// Set up the node
	l.Node.Setup(parent, id)
}
