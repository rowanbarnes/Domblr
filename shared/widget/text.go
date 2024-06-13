package widget

type P struct {
	Node
	Text string
}

func (p *P) Setup(parent *Node, id int) {
	p.Node = Node{
		Structure: Structure{
			Tag: "p",
		},
		Style: &Style{
			Constraint: Constraint{
				Width:  SHRINK,
				Height: SHRINK,
			},
		},
	}
	p.Node.Setup(parent, id)
}
