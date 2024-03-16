package piscine

func AppendRange(min, max int) []int {
	a := []int{}
	if min < max {
		for i := min; i < max; i++ {
			a = append(a, i)
		}
	} else {
		return nil
	}
	return a
}
