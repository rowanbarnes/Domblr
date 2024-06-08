//go:build wasm

package canvas

func clearScreen(width float64, height float64) {
	call("clearRect", 0, 0, width, height)
}

func stroke(color string) {
	set("strokeStyle", color)
}

func fill(color string) {
	set("fillStyle", color)
}

func strokeWidth(width float64) {
	set("lineWidth", width)
}

func rect(x, y, width, height float64) {
	call("fillRect", x, y, width, height)
}
