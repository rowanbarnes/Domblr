package widget

import (
	"bytes"
)

type BuildContext struct {
	Addr       string
	ResPath    string
	Variables  map[int]string
	Constraint Constraint
	ID         int
}

func NewBuildContext(addr string, resPath string) *BuildContext {
	return &BuildContext{
		Addr:       addr,
		ResPath:    resPath,
		Variables:  make(map[int]string),
		Constraint: Constraint{},
		ID:         0,
	}
}

type RenderContext struct {
	Buffer *bytes.Buffer
	CSS    *bytes.Buffer
	HTML   *bytes.Buffer
}

func NewRenderContext() *RenderContext {
	return &RenderContext{
		Buffer: bytes.NewBuffer(make([]byte, 0, 0)),
		CSS:    bytes.NewBuffer(make([]byte, 0, 0)),
		HTML:   bytes.NewBuffer(make([]byte, 0, 0)),
	}
}
