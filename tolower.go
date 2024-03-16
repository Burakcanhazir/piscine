package piscine

func ToLower(s string) string {
	i := []rune(s)
	for a, b := range i {
		if b >= 'A' && b <= 'Z' {
			i[a] = b + ('a' - 'A')
		}
	}
	return string(i)
}
