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
	MinWidth  int
	MaxWidth  int
	MinHeight int
	MaxHeight int
	Style     map[int]string
}

func (s *Style) Design(buffer *bytes.Buffer, id int, pseudo string,
	transform map[int]string, specific map[string]string) {
	// Write the signature `.s$id:$pseudo{`
	buffer.WriteString(".s")
	buffer.WriteString(strconv.Itoa(id))
	if pseudo != "" {
		buffer.WriteString(":")
		buffer.WriteString(pseudo)
	}
	buffer.WriteString("{")

	// Write the transformed properties
	for key, property := range transform {
		if value, ok := s.Style[key]; ok {
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

func Inherit(c **Style, parent *Style) {
	// Handle nil Style
	if *c == nil {
		*c = parent
	} else {
		// Populate size styles
		// TODO (*c).M
	}

	// Copy over the style attributes that the parent has but not the child
	for i, s := range parent.Style {
		if _, ok := (*c).Style[i]; !ok {
			(*c).Style[i] = s
		}
	}
}
