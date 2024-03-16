package piscine

func ConcatParams(args []string) string {
	leng := len(args)
	var str string
	for i := 0; i < leng; i++ {
		if i != leng-1 {
			str = str + args[i] + "\n"
		} else {
			str = str + args[i]
		}
	}
	return str
}
