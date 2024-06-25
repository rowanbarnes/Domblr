package widget

import (
	"Domblr/util"
	"errors"
)

// Options for Axis
const (
	ROW = iota
	COL
)

// Options for Type
const (
	DIV = iota
	UL
)

// Options for justify-contents
const (
	START = iota
	CENTER
	END
)

type List struct {
	Node
	Items       []Widget
	ItemCount   int
	ItemBuilder func(int) Widget
	Axis        int
	Type        int
	// Style contains variables for setting the look of widgets
	// Nullable after Build
	Style map[int]string
}

type listItem struct {
	Node
	Body Widget
}

func (li *listItem) Build(ctx *BuildContext) error {
	// Initialize the Node
	li.Node = Node{
		Structure: Structure{
			Tag: "li",
		},
		Children: []Widget{li.Body},
	}
	util.Panic(li.Node.Build(nil))

	return nil
}

func (l *List) Build(ctx *BuildContext) error {
	// Ensure ItemBuilder is initialized
	if l.ItemBuilder == nil && (l.Items == nil || len(l.Items) < l.ItemCount) {
		return errors.New("list incorrectly declared: ItemBuilder is nil")
	}

	// Construct the remaining children and add them to Items
	if l.Items == nil {
		l.Items = make([]Widget, l.ItemCount)
	}
	for i := len(l.Items); i < l.ItemCount; i++ {
		l.Items = append(l.Items, l.ItemBuilder(i))
	}

	// Construct the remaining children and add them to Items
	for i := 0; i < l.ItemCount; i++ {
		if l.Type == DIV {
			l.Items = append(l.Items, l.ItemBuilder(i))
		} else {
			l.Items = append(l.Items, &listItem{Body: l.ItemBuilder(i)})
		}
	}

	// Initialize Node
	l.Node = Node{
		Structure: Structure{
			Tag: util.If(l.Type == DIV, "div", "ul"),
		},
		Style: Style{
			Properties: map[string]map[string]any{
				"": {
					"background-color": Background,
					//"justify-content":  "space-between", // TODO add some field to switch
					//"align-items":    "center",
					"flex-direction": util.If(l.Axis == ROW, "row", "column"),
				},
			},
			Constraint: Constraint{
				Width:  util.If(l.Axis == ROW, EXPAND, SHRINK),
				Height: util.If(l.Axis == COL, EXPAND, SHRINK),
			},
			Variables: l.Style,
		},
		Children: l.Items,
	}
	util.Panic(l.Node.Build(ctx))

	return nil
}
