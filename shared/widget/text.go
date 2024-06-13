package widget

type P struct {
	Node
	Text string
}

func (p *P) Setup(parent *Node, id int) {
	p.Node = Node{
		Structure: Structure{
			Tag:   "p",
			Inner: p.Text,
		},
	}
	p.Node.Setup(parent, id)
}
