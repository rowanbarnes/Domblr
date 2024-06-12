package widget

import (
	"bytes"
)

type P struct {
	Text  string
	Style *Style
}

func (p *P) Setup(style *Style) {
	if p.Style == nil {
		p.Style = style
	}
}

func (p *P) Design(buffer *bytes.Buffer) *bytes.Buffer {
	return buffer
}

func (p *P) Render(buffer *bytes.Buffer) *bytes.Buffer {
	buffer.WriteString(`<p>`)
	buffer.WriteString(p.Text)
	buffer.WriteString(`</p>`)
	return buffer
}
