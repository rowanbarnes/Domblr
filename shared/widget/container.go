package widget

import (
	"Domblr/shared/comm"
	"bytes"
)

type Container struct {
	Body  Widget
	Style *Style
	id    int
}

func (container *Container) Setup(style *Style) {
	Inherit(&container.Style, style)
	container.id = comm.RegisterElement()

	Setup(container.Body, container.Style)
}

func (container *Container) Design(buffer *bytes.Buffer) *bytes.Buffer {
	container.Style.Design(buffer, container.id, "",
		map[int]string{
			Background: "background-color",
		}, map[string]string{
			"flex": "auto",
		},
	)

	Design(container.Body, buffer)
	return buffer
}

func (container *Container) Render(buffer *bytes.Buffer) *bytes.Buffer {
	OpenTag(buffer, "div", "", container.id, false)
	Render(container.Body, buffer)
	CloseTag(buffer, "div")

	return buffer
}
