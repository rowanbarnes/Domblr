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
	// Style contains variables for setting the look of widgets
	// Nullable after Setup
	Style map[int]string
}

func (l *List) Setup(parent *widget.Node, id int) {
	// Ensure ItemBuilder is initialized
	if l.ItemBuilder == nil {
		// TODO: throw error
		println("List setup error: ItemBuilder is nil")
		return
	}

	// Construct the remaining children and add them to Items
	var items = make([]widget.Widget, l.ItemCount)
	for i := 0; i < l.ItemCount; i++ {
		items = append(items, l.ItemBuilder(i))
	}

	// Initialize Node
	l.Node = widget.Node{
		Structure: widget.Structure{
			Tag: "div",
		},
		Style: widget.Style{
			Properties: map[string]map[string]any{
				"": {
					"background-color": widget.Background,
					//"justify-content":  "space-between", // TODO add some field to switch
					//"align-items":    "center",
					"flex-direction": util.If(l.Axis == ROW, "row", "column"),
				},
			},
			Constraint: widget.Constraint{
				Width:  util.If(l.Axis == ROW, widget.EXPAND, widget.SHRINK),
				Height: util.If(l.Axis == COL, widget.EXPAND, widget.SHRINK),
			},
			Variables: l.Style,
		},
		Children: items,
	}
	l.Node.Setup(parent, id)
}
