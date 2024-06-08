//go:build wasm

package canvas

import (
	"Domblr/shared/comm"
	"Domblr/shared/widget"
	"bytes"
	"math/rand"
	"strconv"
	"syscall/js"
)

type Canvas struct {
	element js.Value
	context js.Value
	id      int
}

func (canvas *Canvas) Setup(_ *widget.Style) {
	canvas.id = comm.RegisterElement()
	println("Getting canvas id: ", canvas.id)
	canvas.element = get("document").Call("getElementById", strconv.Itoa(canvas.id))
	canvas.context = canvas.element.Call("getContext", "2d")
	println("Got canvas context")

	// Setup render loop
	var renderJSCallback js.Func
	renderJSCallback = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		canvas.Draw(get("innerWidth").Float(), get("innerHeight").Float())
		canvas.Update()
		call("setTimeout", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			call("requestAnimationFrame", renderJSCallback)
			return nil
		}), 16)

		return nil
	})
	call("requestAnimationFrame", renderJSCallback)
}

func (canvas *Canvas) Render(buffer *bytes.Buffer) *bytes.Buffer {
	buffer.WriteString("<canvas id=\"")
	buffer.WriteString(strconv.Itoa(canvas.id))
	buffer.WriteString("\"></canvas>")
	return buffer
}

func (canvas *Canvas) Draw(width float64, height float64) {
	clearScreen(width, height)
	fill("red")
	rect(float64(rand.Intn(100)), float64(rand.Intn(100)), 100, 100)
}

func (canvas *Canvas) Update() {}

func call(m string, args ...any) js.Value {
	return js.Global().Call(m, args)
}

func set(p string, x any) {
	js.Global().Set(p, x)
}

func get(p string) js.Value {
	return js.Global().Get(p)
}
