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
	// Style contains variables for setting the look of widgets
	// Nullable after Setup
	Style map[int]string
}

func (l *FlexList) Setup(parent *widget.Node, id int) {
	// Ensure ItemBuilder is initialized
	if l.ItemBuilder == nil && (l.Items == nil || len(l.Items) < l.ItemCount) {
		// TODO: throw error
		println("FlexList setup error: ItemBuilder is nil")
		return
	}

	// Construct the remaining children and add them to Items
	if l.Items == nil {
		l.Items = make([]widget.Widget, l.ItemCount)
	}
	for i := len(l.Items); i < l.ItemCount; i++ {
		l.Items = append(l.Items, l.ItemBuilder(i))
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
		Children: l.Items,
	}
	l.Node.Setup(parent, id)
}
