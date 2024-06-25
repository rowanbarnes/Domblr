package widget

const (
	SHRINK = iota
	EXPAND
)

type Constraint struct {
	Width  int
	Height int
}

func (c *Constraint) Collect(child *Constraint) {
	// Handle nil Constraint
	if child == nil {
		return
	}

	// Collect the child's constraints
	c.Width = max(c.Width, child.Width)
	c.Height = max(c.Height, child.Height)
}
