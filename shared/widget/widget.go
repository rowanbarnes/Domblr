package widget

import (
	"bytes"
)

// Widget interface
type Widget interface {
	// Setup initializes the styles, constraints, structures, etc
	Setup(parent *Node, id int)
	// Render method renders the CSS and HTML code to byte buffers
	Render(css *bytes.Buffer, html *bytes.Buffer)
	// GetConstraint returns the widget's constraints
	GetConstraint() *Constraint
}
