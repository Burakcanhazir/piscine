package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args
	for i := len(a) - 1; i > 0; i-- {
		for _, harf := range []rune(a[i]) {
			z01.PrintRune(harf)
		}
		z01.PrintRune('\n')
	}
}
