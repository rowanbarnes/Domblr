//go:build wasm

package widget

import (
	"Domblr/comm"
	"Domblr/util"
)

type Server struct {
	Addr      string
	ResPath   string
	ApiRouter comm.ApiRouter
}

// RunApp WASM
// 1. `tinygo build -no-debug -o ../../res/main.wasm -target wasm ./main.go`
// 2. `wasm-opt -Oz --shrink-level=2 main.wasm -o main-opt.wasm`
// ~60% size reduction with wasm-opt + gzip
func (server *Server) RunApp(app *App) {
	println("WebAssembly successfully loaded.")

	ctx := NewBuildContext(server.Addr, server.ResPath)
	util.Panic(app.Body.Build(ctx))

	c := make(chan struct{})
	<-c
}
