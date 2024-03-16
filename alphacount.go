package piscine

func AlphaCount(s string) int {
	count := 0
	for _, alicabbar := range s {
		if (alicabbar >= 'A' && alicabbar <= 'Z') || (alicabbar >= 'a' && alicabbar <= 'z') {
			count++
		}
	}
	return count
}
