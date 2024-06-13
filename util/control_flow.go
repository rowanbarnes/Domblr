package util

// If adds a ternary operator, a feature that should be in the language already
// but isn't.
func If[T any](exp bool, trueVal T, falseVal T) T {
	if exp {
		return trueVal
	}
	return falseVal
}
