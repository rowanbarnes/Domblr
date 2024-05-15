//go:build !wasm

package communication

var id int = -1

func onclick(...any) {
}

func RegisterFunc(...any) int {
	id++
	return id
}

func InvokeFunc(...any) {
}

func UpdateWidget(...any) {
}
