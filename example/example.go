package example

import (
	"Domblr/backend"
	"Domblr/frontend"
	"Domblr/shared/comm"
	"Domblr/shared/widget"
	"fmt"
)

func CommunicationApp() *frontend.App {
	return &frontend.App{
		Addr: "http://localhost:8080",
		Page: &widget.Page{
			Config: widget.Config{
				Title: "Backend/Frontend Communication Example",
			},
			Body: &widget.List{
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
			},
		},
	}
}

func CommunicationServer() *backend.Server {
	return &backend.Server{
		Addr: ":8080",
		ApiRouter: comm.ApiRouter{
			"ping": func(params ...any) ([]any, error) {
				return []any{fmt.Sprint("Hello from the server, Button", params[0], "!")}, nil
			},
		},
	}
}
