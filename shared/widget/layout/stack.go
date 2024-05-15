package layout

import (
	"Domblr/shared/widget"
	"bytes"
)

type Stack struct {
	Children []widget.Widget
	Style    *widget.Style
}

func (stack *Stack) Render(buffer *bytes.Buffer) *bytes.Buffer {
	for i := 0; i < len(stack.Children); i++ {
		stack.Children[i].Render(buffer)
	}
	return buffer
}

func (stack *Stack) Setup(style *widget.Style) {
	widget.Inherit(&stack.Style, style)

	for i := 0; i < len(stack.Children); i++ {
		stack.Children[i].Setup(stack.Style)
	}
}
