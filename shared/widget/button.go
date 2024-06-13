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

func (button *Button) Setup(style *Style) {
	if button.Style == nil {
		button.Style = style
	}

	button.id = comm.RegisterFunc(func() {
		button.OnClick(button)
	})
}

func (button *Button) Design(buffer *bytes.Buffer) Constraint {
	// Design the button styling
	button.Style.Design(buffer, button.id, "",
		map[int]string{
			Foreground: "color",
			Padding:    "padding",
		}, map[string]string{
			"text-decoration": "none",
			"position":        "relative",
		},
	)

	// Design hover styles
	button.Style.Design(buffer, button.id, ":hover",
		map[int]string{
			Highlight: "color",
		},
		map[string]string{},
	)

	// Design underline styling
	button.Style.Design(buffer, button.id, "::before",
		map[int]string{
			Highlight: "background-color",
		},
		map[string]string{
			"content":    "\"\"",
			"position":   "absolute",
			"bottom":     "0",
			"left":       "0",
			"width":      "100%",
			"height":     "2px",
			"transform":  "scaleX(0)",
			"transition": "transform 0.2s ease",
		},
	)

	// Design underline hover style
	button.Style.Design(buffer, button.id, ":hover::before",
		map[int]string{},
		map[string]string{
			"transform": "scaleX(1)",
		},
	)
	button.Style.Design(buffer, button.id, ":hover::before",
		map[int]string{},
		map[string]string{
			"transform": "scaleX(1)",
		},
	)

	return Constraint{
		Width:  SHRINK,
		Height: SHRINK,
	}
}

func (button *Button) Render(buffer *bytes.Buffer) {
	OpenTag(buffer, "a", "#", button.id, true)
	buffer.WriteString(button.Label)
	CloseTag(buffer, "a")
	return
}

func (button *Button) SetLabel(Label string) {
	button.Label = Label
	var buffer bytes.Buffer
	button.Render(&buffer)
	comm.UpdateWidget(button.id, buffer.String())
}
