package piscine

func ToUpper(s string) string {
	i := []rune(s)
	for a, b := range i {
		if b >= 'a' && b <= 'z' {
			i[a] = b - ('a' - 'A')
		}
	}
	return string(i)
}
