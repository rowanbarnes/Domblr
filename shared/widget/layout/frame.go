package layout

import (
	"Domblr/shared/widget"
	"bytes"
)

// Frame - each pane of the frame needs to be collapsible, resizeable
type Frame struct {
	Top    widget.Widget
	Bottom widget.Widget
	Left   widget.Widget
	Right  widget.Widget
	Body   widget.Widget
	Style  *widget.Style
}

func (frame *Frame) Render(buffer *bytes.Buffer) *bytes.Buffer {
	buffer.WriteString(`<div style="display: flex; flex-direction: column; height: 100%;">`)
	widget.Render(frame.Top, buffer)
	buffer.WriteString(`</div>`)
	buffer.WriteString(`<div style="display: flex; flex-direction: row; flex: 1;">`)
	widget.Render(frame.Left, buffer)
	widget.Render(frame.Body, buffer)
	widget.Render(frame.Right, buffer)
	buffer.WriteString(`</div>`)
	widget.Render(frame.Bottom, buffer)
	buffer.WriteString(`</div>`)
	return buffer
}

func (frame *Frame) Setup(style *widget.Style) {
	widget.Inherit(&frame.Style, style)
	widget.Setup(frame.Top, frame.Style)
	widget.Setup(frame.Bottom, frame.Style)
	widget.Setup(frame.Left, frame.Style)
	widget.Setup(frame.Right, frame.Style)
	widget.Setup(frame.Body, frame.Style)
}
