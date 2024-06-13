package main

import (
	"Domblr/example"
)

// 1. tinygo build -no-debug -o ../../res/main.wasm -target wasm ./main.go
// 2. wasm-opt -Oz --shrink-level=2 main.wasm -o main-opt.wasm
// ~60% size reduction with wasm-opt + gzip
func main() {
	println("WebAssembly successfully loaded.")
	app := example.App()
	// TODO: default styles
	app.Page.Setup(nil)

	c := make(chan struct{})
	<-c
}
