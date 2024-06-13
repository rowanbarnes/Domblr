package widget

import (
	"Domblr/shared/comm"
	"bytes"
)

type Button struct {
	Node
	Label   string
	OnClick func(button *Button)
}

func (b *Button) Setup(parent *Node, id int) {
	// Set up the node
	b.Node = Node{
		Structure: Structure{
			Tag:     "a",
			Href:    "#",
			Onclick: true,
		},
		Style: &Style{
			Constraint: Constraint{
				Width:  SHRINK,
				Height: SHRINK,
			},
			Properties: map[string]map[string]any{
				"": {
					"color":           Foreground,
					"padding":         Padding,
					"text-decoration": "none",
					"position":        "relative",
				},
				":hover": {
					"color": Highlight,
				},
				// Underline
				"::before": {
					"background-color": Highlight,
					"content":          "\"\"",
					"position":         "absolute",
					"bottom":           "0",
					"left":             "0",
					"width":            "100%",
					"height":           "2px",
					"transform":        "scaleX(0)",
					"transition":       "transform 0.2s ease",
				},
				":hover::before": {
					"transform": "scaleX(1)",
				},
			},
		},
	}

	// Register OnClick
	comm.RegisterFunc(id, func() {
		b.OnClick(b)
	})

	// Setup the node
	b.Node.Setup(parent, id)
}

// SetLabel changes the label of the button and renders the widget
func (b *Button) SetLabel(Label string) {
	b.Label = Label
	var css, html bytes.Buffer
	b.Render(&css, &html)

	// TODO: Handle changing css
	comm.UpdateWidget(b.id, html.String())
}
