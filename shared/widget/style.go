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
	// {`pseudo` : {`property` : inherited variable | `value`}}
	Properties map[string]map[string]any
	// inherited style variables
	inherited map[int]string
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

// Render
// TODO optimize the rendering to not create redundant blocks
func (s *Style) Render(css *bytes.Buffer, id int) {
	// Write the custom styles for each pseudo class
	for pseudo, v := range s.Properties {
		// Write the signature `.s$id$pseudo{`
		css.WriteString(".s")
		css.WriteString(util.ItoABase26(id))
		css.WriteString(pseudo)
		css.WriteString("{")

		// Write the custom styles to the pseudo
		for property, value := range v {
			css.WriteString(property)
			css.WriteString(":")

			// Write either the raw value or inherited variable
			switch value.(type) {
			case string:
				css.WriteString(value.(string))
			case int:
				if value, ok := s.inherited[value.(int)]; ok {
					css.WriteString(value)
				}
			}
			css.WriteString(";")
		}

		// Close the block `}`
		css.WriteString("}")
	}

	// Write the constraint styles
	css.WriteString(".s")
	css.WriteString(util.ItoABase26(id))
	css.WriteString("{")
	css.WriteString("width:")
	css.WriteString(util.If(s.Constraint.Width == SHRINK, "fit-content", "100%"))
	css.WriteString(";width:")
	css.WriteString(util.If(s.Constraint.Height == SHRINK, "fit-content", "100%"))
	css.WriteString(";}")
}
