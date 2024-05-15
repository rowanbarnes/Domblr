package widget

const (
	BgColor = iota
	FgColor
	FontSize
)

type Style struct {
	MinWidth  int
	MaxWidth  int
	MinHeight int
	MaxHeight int
	Style     map[int]string
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
