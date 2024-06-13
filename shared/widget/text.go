package widget

import (
	"Domblr/util"
)

type P struct {
	Node
	Text string
}

func (p *P) Setup(parent *Node, id int) error {
	p.Node = Node{
		Structure: Structure{
			Tag:   "p",
			Inner: p.Text,
		},
		Style: Style{
			Properties: map[string]map[string]any{"": {
				"color":   Foreground,
				"padding": Padding,
			}},
		},
	}
	util.Panic(p.Node.Setup(parent, id))

	return nil
}
