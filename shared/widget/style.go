package widget

import (
	"bytes"
	"strconv"
)

const (
	// Colors
	Background = iota
	Foreground
	Highlight
	Alert

	// Fonts
	FontSize

	// Layout
	Padding

	// Misc
	Roundness
)

type Style struct {
	Constraint
	inherited map[int]string

}

func (s *Style) Inherit(parent *Style) {
	// Copy over the style attributes that the parent has but not the child
	for i, p := range parent.inherited {
		if _, ok := s.inherited[i]; !ok {
			s.inherited[i] = p
		}
	}
}

func (s *Style) Render(buffer *bytes.Buffer, id int, pseudo string,
	transform map[int]string, specific map[string]string) {

	// Write the signature `.s$id$pseudo{`
	buffer.WriteString(".s")
	buffer.WriteString(strconv.Itoa(id))
	if pseudo != "" {
		buffer.WriteString(pseudo)
	}
	buffer.WriteString("{")

	// Write the transformed inherited
	for key, property := range transform {
		if value, ok := s.inherited[key]; ok {
			buffer.WriteString(property)
			buffer.WriteString(":")
			buffer.WriteString(value)
			buffer.WriteString(";")
		}
	}

	// Write the specific styles
	for property, value := range specific {
		buffer.WriteString(property)
		buffer.WriteString(":")
		buffer.WriteString(value)
		buffer.WriteString(";")
	}

	// Close the block `}`
	buffer.WriteString("}")
}
