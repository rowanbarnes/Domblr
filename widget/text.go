package widget

import (
	"Domblr/util"
)

type Text struct {
	Node
	Text string
	// Style contains variables for setting the look of widgets
	// Nullable after Build
	Style map[int]string
}

func (t *Text) Build(ctx *BuildContext) error {
	t.Node = Node{
		Structure: Structure{
			Tag:   "p",
			Inner: t.Text,
		},
		Style: Style{
			Properties: map[string]map[string]any{"": {
				"color":   Foreground,
				"padding": Padding,
			}},
			Variables: t.Style,
		},
	}
	util.Panic(t.Node.Build(ctx))

	return nil
}
