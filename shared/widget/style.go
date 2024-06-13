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
	// Constraint composes the layout constraints into Style
	Constraint
	// Properties contains CSS properties in the form:
	// {`pseudo` : {`property` : var (int) | `value`}}
	Properties map[string]map[string]any
	// Variables holds the inherited styles
	Variables map[int]string
}

func (s *Style) Setup(parent *Style) {
	// Initialize Properties
	if s.Properties == nil {
		s.Properties = make(map[string]map[string]any)
	}

	// Initialize Variables
	if s.Variables == nil {
		s.Variables = make(map[int]string)
	}

	// Handle nil parent
	if parent == nil {
		return
	}

	// Inherit style Variables from the parent
	for i, p := range parent.Variables {
		if _, ok := s.Variables[i]; !ok {
			s.Variables[i] = p
		}
	}
}

// Render
// TODO optimize the rendering to not create redundant blocks
func (s *Style) Render(css *bytes.Buffer, id int) {
	// Write the custom styles for each pseudo class
	for pseudo, pseudoProps := range s.Properties {
		// Write the signature `.s$id$pseudo{`
		css.WriteString(".")
		css.WriteString(util.ItoABase26(id))
		css.WriteString(pseudo)
		css.WriteString("{")

		// Write the custom styles to the pseudo
		// Write either the raw string value or the retrieved Variables' variable
		for property, value := range pseudoProps {
			switch value.(type) {
			case string:
				css.WriteString(property)
				css.WriteString(":")
				css.WriteString(value.(string))
				css.WriteString(";")
			case int:
				if varValue, ok := s.Variables[value.(int)]; ok {
					css.WriteString(property)
					css.WriteString(":")
					css.WriteString(varValue)
					css.WriteString(";")
				}
			}
		}

		// Close the block `}`
		css.WriteString("}")
	}

	// Write the constraint styles
	css.WriteString(".")
	css.WriteString(util.ItoABase26(id))
	css.WriteString("{")
	css.WriteString("width:")
	css.WriteString(util.If(s.Constraint.Width == SHRINK, "fit-content", "100%"))
	css.WriteString(";height:")
	css.WriteString(util.If(s.Constraint.Height == SHRINK, "fit-content", "100%"))
	css.WriteString(";}")
}
