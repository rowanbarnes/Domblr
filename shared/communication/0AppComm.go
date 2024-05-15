//go:build wasm

package communication

import (
	"syscall/js"
)

var funcs []func()

//export onclick
func onclick(id int) {
	InvokeFunc(id)
}

func RegisterFunc(fn func()) int {
	funcs = append(funcs, fn)
	return len(funcs) - 1
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
