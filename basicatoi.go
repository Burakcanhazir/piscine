package piscine

func BasicAtoi(s string) int {
	result := 0

	for _, digit := range s {
		// Check if the character is a digit
		if digit >= '0' && digit <= '9' {
			// Convert the digit character to an integer
			// and add it to the result
			result = result*10 + int(digit-'0')
		} else {
			// If any non-digit character is encountered, break the loop
			break
		}
	}

	return result
}
