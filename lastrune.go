package piscine

func LastRune(s string) rune {
	a := []rune(s)
	b := len(a)
	return a[b-1]
}
