package widget

import (
	"bytes"
)

// Widget interface
type Widget interface {
	// Setup initializes the styles, constraints, structures, etc
	Setup(parent *Node, id int) error
	// Render method renders the CSS and HTML code to byte buffers
	Render(css *bytes.Buffer, html *bytes.Buffer)
	// GetConstraint returns the widget's constraints
	GetConstraint() *Constraint
	// GetDescendants returns the number of children in this widget
	GetDescendants() int
}
