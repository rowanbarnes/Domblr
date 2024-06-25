package widget

import (
	"Domblr/util"
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

func (s *Style) Build(variables *map[int]string) {
	// Initialize Properties
	if s.Properties == nil {
		s.Properties = make(map[string]map[string]any)
	}

	// Initialize Variables
	if s.Variables == nil {
		s.Variables = make(map[int]string)
	}

	// Handle nil variables
	if variables == nil {
		return
	}

	// Inherit style Variables from the parent
	for i, p := range *variables {
		if _, ok := s.Variables[i]; !ok {
			s.Variables[i] = p
		}
	}
}

// Render
// TODO optimize the rendering to not create redundant blocks
func (s *Style) Render(ctx *RenderContext, id int) {
	// Write the custom styles for each pseudo class
	for pseudo, pseudoProps := range s.Properties {
		// Write the signature `.s$id$pseudo{`
		ctx.CSS.WriteString(".")
		ctx.CSS.WriteString(util.ItoABase26(id))
		ctx.CSS.WriteString(pseudo)
		ctx.CSS.WriteString("{")

		// Write the custom styles to the pseudo
		// Write either the raw string value or the retrieved Variables' variable
		for property, value := range pseudoProps {
			switch value.(type) {
			case string:
				ctx.CSS.WriteString(property)
				ctx.CSS.WriteString(":")
				ctx.CSS.WriteString(value.(string))
				ctx.CSS.WriteString(";")
			case int:
				if varValue, ok := s.Variables[value.(int)]; ok {
					ctx.CSS.WriteString(property)
					ctx.CSS.WriteString(":")
					ctx.CSS.WriteString(varValue)
					ctx.CSS.WriteString(";")
				}
			}
		}

		// Close the block `}`
		ctx.CSS.WriteString("}")
	}

	// Write the constraint styles
	ctx.CSS.WriteString(".")
	ctx.CSS.WriteString(util.ItoABase26(id))
	ctx.CSS.WriteString("{")
	ctx.CSS.WriteString("width:")
	ctx.CSS.WriteString(util.If(s.Constraint.Width == SHRINK, "fit-content", "100%"))
	ctx.CSS.WriteString(";height:")
	ctx.CSS.WriteString(util.If(s.Constraint.Height == SHRINK, "fit-content", "100%"))
	ctx.CSS.WriteString(";}")
}
