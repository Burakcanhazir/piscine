package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	for i := 0; i <= len(s)-len(toFind); i++ {
		x := true
		for j := 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] && s[i+j] != toFind[j]+32 && s[i+j] != toFind[j]-32 {
				x = false
				break
			}
		}
		if x {
			return i
		}
	}
	return -1
}
