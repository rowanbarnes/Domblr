package widget

type Container struct {
	Node
	// Body is the contents of the container
	// Nullable after Setup
	Body Widget
	// Style contains Variables for setting the look of widgets
	// Nullable after Setup
	Style map[int]string
}

func (c *Container) Setup(parent *Node, id int) {
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
	c.Node.Setup(parent, id)
}
