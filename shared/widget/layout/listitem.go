package layout

import (
	"Domblr/shared/widget"
)

type listItem struct {
	widget.Node
	Body widget.Widget
}

func (li *listItem) Setup(parent *widget.Node, id int) {
	// Initialize the Node
	li.Node = widget.Node{
		Structure: widget.Structure{
			Tag: "li",
		},
		Style: widget.Style{
			Constraint: widget.Constraint{
				Width:  widget.SHRINK,
				Height: widget.SHRINK,
			},
		},
		Children: []widget.Widget{li.Body},
	}
	li.Node.Setup(parent, id)
}
