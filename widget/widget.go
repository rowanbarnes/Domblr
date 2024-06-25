package widget

// Widget interface
type Widget interface {
	// Build initializes the styles, constraints, structures, etc
	Build(ctx *BuildContext) error
	// Render method renders the CSS and HTML code to byte buffers
	Render(ctx *RenderContext)
}
