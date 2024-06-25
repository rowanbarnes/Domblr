package widget

import (
	"Domblr/util"
	"slices"
)

// Node represents an abstract node in the DOM app
type Node struct {
	// Structure represents the HTML structure of this node
	Structure
	// Style represents the CSS styling of this node
	Style
	// ID the identifier of this node, used for ids and classes
	ID int
	// Children holds the widgets that have been added to this node and setup
	Children []Widget
}

// Build sets the nodes' ID and inherits style properties from its parent.
// Consider switching to the `NewNode` pattern instead
func (n *Node) Build(ctx *BuildContext) error {
	n.ID = ctx.ID
	ctx.ID++
	n.Style.Build(&ctx.Variables)

	// Initialize Children
	if n.Children == nil {
		n.Children = make([]Widget, 0)
	} else {
		// Remove nil children
		n.Children = slices.DeleteFunc(n.Children, func(widget Widget) bool {
			return widget == nil
		})
	}

	// Build children
	for _, c := range n.Children {
		ctx.Variables = n.Style.Variables
		util.Panic(c.Build(ctx))
		n.Collect(&ctx.Constraint)
	}
	ctx.Constraint = n.Constraint

	return nil
}

// Render outputs CSS and HTML code to the given relevant buffers.
func (n *Node) Render(ctx *RenderContext) {
	n.Style.Render(ctx, n.ID)
	n.Structure.Render(ctx, n)
}
