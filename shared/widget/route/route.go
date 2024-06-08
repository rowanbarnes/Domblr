package route

import (
	"Domblr/shared/widget"
	"bytes"
)

type Route struct {
	ID    string
	Body  widget.Widget
	Style *widget.Style
}

func (route *Route) Setup(style *widget.Style) {
	if route.Style == nil {
		route.Style = style
	}
	route.Body.Setup(route.Style)
}

func (route *Route) Render(buffer *bytes.Buffer) *bytes.Buffer {
	// if

	route.Body.Render(buffer)

	return buffer
}
