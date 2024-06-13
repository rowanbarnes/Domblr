package widget

import (
	"bytes"
	"slices"
)

// Node represents an abstract node in the DOM tree
type Node struct {
	// Structure represents the HTML structure of this node
	Structure
	// Style represents the CSS styling of this node
	Style
	// Children holds the widgets that have been added to this node and setup
	Children []Widget
	// id the identifier of this node, used for ids and classes
	id int
	// descendants is the number of nodes in the subtree rooted at this Node
	descendants int
}

// Setup sets the nodes' id and inherits style properties from its parent.
func (n *Node) Setup(parent *Node, id int) {
	n.id = id
	n.descendants = 1
	if parent != nil {
		n.Style.Setup(&parent.Style)
	}

	// Initialize Children
	if n.Children == nil {
		n.Children = make([]Widget, 0)
	}

	// Remove nil children
	n.Children = slices.DeleteFunc(n.Children, func(widget Widget) bool {
		return widget == nil
	})

	// Setup children
	for _, c := range n.Children {
		c.Setup(n, n.id+n.descendants)
		n.descendants += c.GetDescendants()
		n.Collect(c.GetConstraint())
	}
}

// Render outputs CSS and HTML code to the given relevant buffers.
func (n *Node) Render(css *bytes.Buffer, html *bytes.Buffer) {
	n.Style.Render(css, n.id)
	n.Structure.Render(css, html, n)
}

// GetConstraint returns this nodes' constraint
func (n *Node) GetConstraint() *Constraint {
	return &n.Style.Constraint
}

func (n *Node) GetDescendants() int {
	return n.descendants
}
