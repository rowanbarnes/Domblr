//go:build !wasm

package comm

var id = -1

func onclick(_ ...any) {
}

func RegisterFunc(_ ...any) int {
	id++
	return id
}

func RegisterElement() int {
	id++
	return id
}

func InvokeFunc(_ ...any) {
}

func UpdateWidget(_ ...any) {
}
