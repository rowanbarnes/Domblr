//go:build !wasm

package canvas

import (
	"Domblr/shared/comm"
	"Domblr/shared/widget"
	"bytes"
	"strconv"
)

type Canvas struct {
	id int
}

func (canvas *Canvas) Setup(style *widget.Style) {
	canvas.id = comm.RegisterElement()
}

func (canvas *Canvas) Render(buffer *bytes.Buffer) *bytes.Buffer {
	buffer.WriteString("<canvas id=\"")
	buffer.WriteString(strconv.Itoa(canvas.id))
	buffer.WriteString("\"></canvas>")
	return buffer
}

func (canvas *Canvas) Draw(_ float64, _ float64) {}

func (canvas *Canvas) Update() {}
