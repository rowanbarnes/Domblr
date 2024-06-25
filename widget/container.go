package widget

import (
	"Domblr/util"
)

type Container struct {
	Node
	// Body is the contents of the container
	// Nullable after Build
	Body Widget
	// Style contains Variables for setting the look of widgets
	// Nullable after Build
	Style map[int]string
}

func (c *Container) Build(ctx *BuildContext) error {
	// Initialize the Node
	c.Node = Node{
		Structure: Structure{
			Tag: "div",
		},
		Style: Style{
			Properties: map[string]map[string]any{
				"": {
					"background-color": Background,
				},
			},
			Variables: c.Style,
		},
		Children: []Widget{c.Body},
	}
	util.Panic(c.Node.Build(nil))

	return nil
}
