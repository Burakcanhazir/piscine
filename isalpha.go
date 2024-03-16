package piscine

func IsAlpha(s string) bool {
	for _, x := range s {
		if x < '0' || x > '9' && x < 'A' || x > 'Z' && x < 'a' || x > 'z' {
			return false
		}
	}
	return true
}
