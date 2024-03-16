package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var compt [10]int
	for n != 0 {
		compt[n%10]++
		n /= 10
	}
	for num := 0; num < 10; num++ {
		for compt[num] > 0 {
			z01.PrintRune(rune(num) + '0')
			compt[num]--
		}
	}
}
