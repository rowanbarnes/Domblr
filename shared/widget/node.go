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
	// TODO consider removing constraint
	// The same behaviour could be achieved without storing the constraint
	// data. Instead, it could be passed as a pointer up the tree and modified as it
	// goes.
	// Although, it might be handy to have a copy of each nodes' constraints in the
	// future.
	*Constraint
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
	n.Style.Render(css, n.id, "", map[int]string{}, map[string]string{})
	n.Structure.Render(html)
	for _, child := range n.children {
		child.Render(css, html)
	}
}

// AddChild adds a child widget to the node.
func (n *Node) AddChild(child ...Widget) {
	n.children = append(n.children, child...)
}

// GetConstraint returns this nodes' constraint
func (n *Node) GetConstraint() *Constraint {
	return n.Constraint
}
