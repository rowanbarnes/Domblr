package widget

import (
	"bytes"
)

// Node represents an abstract node in the DOM tree
type Node struct {
	// Structure represents the HTML structure of this node
	Structure
	// Style represents the CSS styling of this node
	*Style
	// children holds the widgets that have been added to this node and setup
	children []Widget
	// id the identifier of this node, used for ids and classes
	id int
}

// Setup sets the nodes' id and inherits style properties from its parent.
func (n *Node) Setup(parent *Node, id int) {
	n.id = id
	n.Inherit(parent.Style)
	for i, child := range n.children {
		child.Setup(n, id+i)
		n.Collect(child.GetConstraint())
	}
}

// Render outputs CSS and HTML code to the given relevant buffers.
func (n *Node) Render(css *bytes.Buffer, html *bytes.Buffer) {
	n.Style.Render(css, n.id)
	n.Structure.Render(css, html, n)
}

// AddChild adds a child widget to the node.
func (n *Node) AddChild(child ...Widget) {
	n.children = append(n.children, child...)
}

// GetConstraint returns this nodes' constraint
func (n *Node) GetConstraint() *Constraint {
	return &n.Style.Constraint
}
