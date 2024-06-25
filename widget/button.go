package widget

import (
	"Domblr/comm"
	"Domblr/util"
)

type Button struct {
	Node
	Label   string
	OnClick func(button *Button)
	// Style contains Variables for setting the look of widgets
	// Nullable after Build
	Style map[int]string
}

func (b *Button) Build(ctx *BuildContext) error {
	// Initialize Node
	b.Node = Node{
		Structure: Structure{
			Tag:     "a",
			Href:    "#",
			Inner:   b.Label,
			Onclick: true,
		},
		Style: Style{
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
			Variables: b.Style,
		},
	}
	util.Panic(b.Node.Build(ctx))

	// Register OnClick
	comm.RegisterFunc(b.ID, func() {
		b.OnClick(b)
	})
	return nil
}

// SetLabel changes the label of the button and renders the widget
func (b *Button) SetLabel(Label string) {
	b.Label = Label
	b.Structure.Inner = Label
	var ctx = NewRenderContext()
	b.Render(ctx)

	// TODO: Handle changing css
	comm.UpdateWidget(b.ID, ctx.HTML.String())
}
