package widget

import (
	"Domblr/shared/communication"
	"bytes"
	"strconv"
)

type Button struct {
	Label   string
	OnClick func(button *Button)
	id      int
	Style   *Style
}

func (button *Button) Render(buffer *bytes.Buffer) *bytes.Buffer {
	buffer.WriteString("<button id=\"")
	buffer.WriteString(strconv.Itoa(button.id))
	buffer.WriteString("\" onclick=\"wasm.exports.onclick(")
	buffer.WriteString(strconv.Itoa(button.id))
	buffer.WriteString(")\">")
	buffer.WriteString(button.Label)
	buffer.WriteString("</button>")
	return buffer
}

func (button *Button) Setup(style *Style) {
	if button.Style == nil {
		button.Style = style
	}

	button.id = communication.RegisterFunc(func() {
		button.OnClick(button)
	})
}

func (button *Button) SetLabel(Label string) {
	button.Label = Label
	var buffer bytes.Buffer
	button.Render(&buffer)
	communication.UpdateWidget(button.id, buffer.String())
}
