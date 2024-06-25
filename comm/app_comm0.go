//go:build wasm

package comm

import (
	"syscall/js"
)

var funcs = make(map[int]func())

//export invoke
func invoke(id int) {
	funcs[id]()
}

func RegisterFunc(id int, fn func()) {
	funcs[id] = fn
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

func PushState(url string) {
	window := js.Global().Get("window")
	history := window.Get("history")
	history.Call("pushState", "", "", "/"+url)
	println("Pushing state: ", "/"+url)
}
