package widget

import (
	"Domblr/shared/comm"
	"bytes"
)

type Button struct {
	Label   string
	OnClick func(button *Button)
	Style   *Style
	id      int
}

func (button *Button) Setup(style *Style) {
	if button.Style == nil {
		button.Style = style
	}

	button.id = comm.RegisterFunc(func() {
		button.OnClick(button)
	})
}

func (button *Button) Design(buffer *bytes.Buffer) *bytes.Buffer {
	button.Style.Design(buffer, button.id, "",
		map[int]string{
			Foreground: "color",
			Padding:    "padding",
		}, map[string]string{
			"text-decoration": "none",
		},
	)

	button.Style.Design(buffer, button.id, "hover",
		map[int]string{
			Highlight: "color",
		},
		map[string]string{},
	)

	return buffer
}

func (button *Button) Render(buffer *bytes.Buffer) *bytes.Buffer {
	OpenTag(buffer, "a", "#", button.id, true)
	buffer.WriteString(button.Label)
	CloseTag(buffer, "a")
	return buffer
}

func (button *Button) SetLabel(Label string) {
	button.Label = Label
	var buffer bytes.Buffer
	button.Render(&buffer)
	comm.UpdateWidget(button.id, buffer.String())
}
