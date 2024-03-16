package piscine

func IsLower(s string) bool {
	for _, charNumber := range s {
		if charNumber < 'a' || charNumber > 'z' {
			return false
		}
	}
	return true
}
