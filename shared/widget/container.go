package widget

import (
	"bytes"
)

type Container struct {
	Body  Widget
	Style *Style
}

func (container *Container) Render(buffer *bytes.Buffer) *bytes.Buffer {
	buffer.WriteString("<div style=\"background-color: " + container.Style.Style[BgColor] + "\">")
	Render(container.Body, buffer)
	buffer.WriteString("</div>")
	return buffer
}

func (container *Container) Setup(style *Style) {
	Inherit(&container.Style, style)
	Setup(container.Body, container.Style)
}
