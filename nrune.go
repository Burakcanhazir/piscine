package piscine

func NRune(s string, n int) rune {
	i := []rune(s)

	if n < 1 || n > len(i) {
		return 0
	}
	return i[n-1]
}
