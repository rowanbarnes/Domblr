package util

// Basic control flow utilities to make the language less verbose and annoying

// If acts as a ternary operator
func If[T any](exp bool, trueVal T, falseVal T) T {
	if exp {
		return trueVal
	}
	return falseVal
}

// Panic will panic and exit the program if it is passed a non-nil error
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
