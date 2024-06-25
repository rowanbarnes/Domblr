package widget

import (
	"Domblr/comm"
	"Domblr/util"
)

type Image struct {
	Node
	Path string
	// Style contains variables for setting the look of widgets
	// Nullable after Build
	Style map[int]string
}

func (i *Image) Build(ctx *BuildContext) error {
	i.Node = Node{
		Structure: Structure{
			Tag: "img",
			Src: "data:image/jpeg;base64,",
		},
		Style: Style{
			Properties: map[string]map[string]any{"": {
				"padding": Padding,
			}},
			Variables: i.Style,
			Constraint: Constraint{
				Width:  EXPAND,
				Height: EXPAND,
			},
		},
	}
	util.Panic(i.Node.Build(ctx))

	comm.LoadImage(ctx.ResPath + i.Path).Then(func(param []any) {
		println(".Then")
		src, ok := param[0].(string)
		if !ok {
			println("Not OK...")
		}
		i.Node.Structure.Src = src
		println("Image src: " + i.Node.Structure.Src)
		var ctx = NewRenderContext()
		i.Render(ctx)
		comm.UpdateWidget(i.ID, ctx.HTML.String())
	})

	return nil
}
