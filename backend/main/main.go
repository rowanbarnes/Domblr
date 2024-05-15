package main

import "Domblr/example"

func main() {
	app := example.App()
	app.Page.Setup(nil)
	svr := example.Server()
	svr.Serve(app)
}
