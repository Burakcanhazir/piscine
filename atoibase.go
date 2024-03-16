package piscine

func AtoiBase(s string, base string) int {
	basement := []rune(base)
	str := []rune(s)
	baseNumber := len(basement)

	result := 0
	isNegative := false

	if str[0] == '-' {
		isNegative = true
	}
	v := true
	if baseNumber < 2 {
		v = false
	} else {
		for i := 0; i < baseNumber; i++ {
			if basement[i] == '-' || basement[i] == '+' {
				v = false
			}
			for j := i + 1; j < baseNumber; j++ {
				if basement[i] == basement[j] {
					v = false
				}
			}
		}
	}
	if !v {
		return 0
	}
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(basement); j++ {
			if basement[j] == str[i] {
				result = baseNumber*result + j
				break
			}
		}
	}
	if isNegative {
		result *= -1
	}
	return result
}
