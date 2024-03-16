package piscine

func BasicAtoi2(s string) int {
	result := 0

	for _, digit := range s {
		if digit < '0' || digit > '9' {
			return 0
			// Check if the character is a digit
		} else if digit >= '0' && digit <= '9' {
			// Convert the digit character to an integer
			// and add it to the result
			result = result*10 + int(digit-'0')
		} else if digit == ' ' {
			return 0
		} else {
			// If any non-digit character is encountered, break the loop
			break
		}
	}

	return result
}
