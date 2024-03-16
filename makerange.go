package piscine

func MakeRange(min, max int) []int {
	a := []int{}
	if min < max {
		a = make([]int, max-min)
		for i := 0; i < max-min; i++ {
			a[i] = i + min
		}
	} else {
		return nil
	}
	return a
}
