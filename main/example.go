package main

import (
	"Domblr/comm"
	"Domblr/widget"
	"fmt"
)

func main() {
	(&widget.Server{
		Addr:    ":8080",
		ResPath: "./res/",
		ApiRouter: comm.ApiRouter{
			"ping": func(params ...any) ([]any, error) {
				return []any{fmt.Sprint("Hello from the server, Button", params[0], "!")}, nil
			},
		},
	}).RunApp(&widget.App{
		Config: widget.Config{
			Title: "Backend/Frontend Communication Example",
		},
		Body: &widget.List{
			Axis: widget.COL,
			Items: []widget.Widget{
				&widget.List{
					ItemCount: 5,
					ItemBuilder: func(i int) widget.Widget {
						return &widget.Button{
							Label: fmt.Sprint("Button", i),
							OnClick: func(button *widget.Button) {
								comm.Call("ping", i).Then(func(res []any) {
									button.SetLabel(res[0].(string))
								})
							},
						}
					},
				}, &widget.Image{Path: "image.png"},
			},
		},
	})
}
