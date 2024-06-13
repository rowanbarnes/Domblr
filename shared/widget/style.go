package widget

import (
	"Domblr/util"
	"bytes"
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
	Transform map[int]string
	Custom    map[string]map[string]string
}

func (s *Style) Inherit(parent *Style) {
	// Handle nil parent
	if parent == nil {
		return
	}

	// Copy over the style attributes that the parent has but not the child
	for i, p := range parent.inherited {
		if _, ok := s.inherited[i]; !ok {
			s.inherited[i] = p
		}
	}
}

func (s *Style) Render(css *bytes.Buffer, id int) {
	// Write the signature `.s$id{`
	css.WriteString(".s")
	css.WriteString(util.ItoABase26(id))
	css.WriteString("{")

	// Write the transformed inherited styles, only to the base style class
	for key, property := range s.Transform {
		if value, ok := s.inherited[key]; ok {
			css.WriteString(property)
			css.WriteString(":")
			css.WriteString(value)
			css.WriteString(";")
		}
	}

	// Close the block `}`
	css.WriteString("}")

	// Write the custom styles for each pseudo class
	for pseudo, properties := range s.Custom {
		// Write the signature `.s$id$pseudo{`
		css.WriteString(".s")
		css.WriteString(util.ItoABase26(id))
		css.WriteString(pseudo)
		css.WriteString("{")

		// Write the custom styles to the pseudo
		for property, value := range properties {
			css.WriteString(property)
			css.WriteString(":")
			css.WriteString(value)
			css.WriteString(";")
		}

		// Close the block `}`
		css.WriteString("}")
	}
}
