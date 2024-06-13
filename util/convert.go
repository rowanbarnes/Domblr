package util

// ItoABase26 converts a number to a string containing lowercase letters.
func ItoABase26(n int) string {
	if n < 0 {
		// TODO consider throwing an error
		return "invalid"
	}

	result := ""

	for n >= 0 {
		remainder := n % 26
		result = string(rune('a'+remainder)) + result
		n /= 26
		n--
	}

	return result
}
