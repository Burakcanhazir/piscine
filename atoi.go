package piscine

func Atoi(s string) int {
	result := 0
	sign := 1
	started := false

	for _, char := range s {
		if char == ' ' && !started {
			continue
		} else if char == '+' && !started {
			started = true
		} else if char == '-' && !started {
			started = true
			sign = -1
		} else if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
			started = true
		} else {
			return 0 // Invalid character encountered
		}
	}

	return result * sign
}
