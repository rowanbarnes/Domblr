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
	}
	util.Panic(p.Node.Setup(parent, id))

	return nil
}
