package piscine

func Capitalize(s string) string {
	result := []rune(s)
	capitalizeNext := true

	for i, char := range result {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if capitalizeNext {
				result[i] = toUpperCase(char)
			} else {
				result[i] = toLowerCase(char)
			}
			capitalizeNext = false
		} else {
			capitalizeNext = true
		}
	}

	return string(result)
}

func toUpperCase(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return char - ('a' - 'A')
	}
	return char
}

func toLowerCase(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + ('a' - 'A')
	}
	return char
}
