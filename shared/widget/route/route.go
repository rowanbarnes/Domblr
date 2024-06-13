package route

import (
	"Domblr/shared/comm"
	"Domblr/shared/widget"
	"Domblr/util"
	"bytes"
)

type Route struct {
	widget.Node
	// Router is a mapping from string identifiers to Widgets
	Router map[string]widget.Widget
	// Style contains Variables for setting the look of widgets
	// Nullable after Setup
	Style map[int]string
}

func (r *Route) Setup(parent *widget.Node, id int) error {
	// Initialize router
	if r.Router == nil {
		r.Router = make(map[string]widget.Widget)
	}

	// Initialize node
	r.Node = widget.Node{
		Structure: widget.Structure{
			Tag: "div",
		},
		Style: widget.Style{
			Variables: r.Style,
		},
		Children: []widget.Widget{r.Router[""]},
	}
	util.Panic(r.Node.Setup(parent, id))

	return nil
}

func (r *Route) ChangeRoute(id string) {
	if _, ok := r.Router[id]; !ok {
		return
	}
	println("Route found")

	r.Node.Children = []widget.Widget{r.Router[id]}
	util.Panic(r.Node.Setup(&r.Node, r.ID))
	var css, html bytes.Buffer
	r.Render(&css, &html)

	//TODO: handle changing css
	comm.UpdateWidget(r.ID, html.String())
}
