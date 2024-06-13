package widget

type Container struct {
	Node
	Body Widget
}

func (c *Container) Setup(parent *Node, id int) {
	c.Node = Node{
		Structure: Structure{
			Tag: "div",
		},
		Style: &Style{
			Constraint: Constraint{
				Width:  SHRINK,
				Height: SHRINK,
			},
			Properties: map[string]map[string]any{
				"": {
					"background-color": Background,
				},
			},
		},
	}

	c.Node.Setup(parent, id)
}
