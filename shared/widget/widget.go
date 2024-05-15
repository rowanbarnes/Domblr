package widget

import (
	"bytes"
)

// Widget interface
type Widget interface {
	// Setup populates the widget tree
	Setup(style *Style)
	// Render method returns HTML string
	Render(buffer *bytes.Buffer) *bytes.Buffer
}

func Setup(widget Widget, style *Style) {
	if widget != nil {
		widget.Setup(style)
	}
}

func Render(widget Widget, buffer *bytes.Buffer) {
	if widget != nil {
		widget.Render(buffer)
	}
}
