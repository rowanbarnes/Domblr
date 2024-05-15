package example

import (
	"Domblr/backend"
	"Domblr/frontend"
	"Domblr/shared/communication"
	"Domblr/shared/widget"
	"Domblr/shared/widget/layout"
	"fmt"
)

func App() *frontend.App {
	return &frontend.App{
		Addr: "http://localhost:8080",
		Page: &widget.Page{
			Title: "Hello World Title",
			Style: &widget.Style{},
			Body: &layout.Frame{
				Top: &layout.List{
					ItemCount: 5,
					ItemBuilder: func(i int) widget.Widget {
						return &widget.Button{Label: fmt.Sprint("Button", i)}
					},
				},
				Left: &layout.List{
					ItemCount: 10,
					ItemBuilder: func(i int) widget.Widget {
						return &widget.Container{
							//Style: &widget.Style{
							//	BgColor: fmt.Sprint("rgba(", i*25, ", 0, 0, 1.0)"),
							//},
							Body: &widget.P{
								Text: "text",
							},
						}
					},
				},
				Body: &widget.Container{
					//Style: &widget.Style{
					//	BgColor: fmt.Sprint("rgba(255, 0, 0, 1.0)"),
					//},
				},
			},
		},
	}
}

func Server() *backend.Server {
	return &backend.Server{
		Addr: ":8080",
		ApiRouter: communication.ApiRouter{
			"hello": func(params ...any) ([]any, error) {
				return []any{"Hello, world!"}, nil
			},
		},
	}
}
