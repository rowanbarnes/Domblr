package widget

import (
	"Domblr/comm"
	"Domblr/util"
)

type State struct {
	Node
	// Router is a mapping from string identifiers to Widgets
	Router map[string]Widget
	// Style contains Variables for setting the look of widgets
	// Nullable after Build
	Style map[int]string
}

func (r *State) Build(ctx *BuildContext) error {
	// Initialize router
	if r.Router == nil {
		r.Router = make(map[string]Widget)
	}

	// Initialize node
	r.Node = Node{
		Structure: Structure{
			Tag: "div",
		},
		Style: Style{
			Variables: r.Style,
		},
		Children: []Widget{r.Router[""]},
	}
	util.Panic(r.Node.Build(nil))

	return nil
}

func (r *State) ChangeRoute(id string) {
	if _, ok := r.Router[id]; !ok {
		return
	}
	println("State found")

	r.Node.Children = []Widget{r.Router[id]}
	util.Panic(r.Node.Build(nil))
	ctx := NewRenderContext()
	r.Render(ctx)

	//TODO: handle changing css
	println("UpdateWidget")
	comm.UpdateWidget(r.ID, ctx.HTML.String())
	println("PushState")
	comm.PushState(id)
	println("Pushed state")
}
