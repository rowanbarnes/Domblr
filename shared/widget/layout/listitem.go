package layout

import (
	"Domblr/shared/widget"
)

type ListItem struct {
	widget.Node
	Body widget.Widget
}

func (li *ListItem) Setup(parent *widget.Node, id int) {
	li.Node = widget.Node{
		Structure: widget.Structure{
			Tag: "li",
		},
		Style: &widget.Style{
			Constraint: widget.Constraint{
				Width:  widget.SHRINK,
				Height: widget.SHRINK,
			},
		},
	}

	li.Node.Setup(parent, id)
}
