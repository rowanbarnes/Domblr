//go:build wasm

package comm

import (
	"syscall/js"
)

var funcs = make(map[int]func())
var id = -1

//export onclick
func onclick(id int) {
	InvokeFunc(id)
}

func RegisterFunc(fn func()) int {
	id++
	funcs[id] = fn
	return id
}

func RegisterElement() int {
	id++
	return id
}

func InvokeFunc(id int) {
	funcs[id]()
}

func UpdateWidget(id int, html string) {
	document := js.Global().Get("document")
	element := document.Call("getElementById", id)
	if element.Truthy() {
		element.Set("outerHTML", html)
	} else {
		js.Global().Get("console").Call("error",
			"Element with ID not found: ", js.ValueOf(id).String())
	}
}
