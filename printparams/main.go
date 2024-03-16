package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args
	for i := 1; i < len(a); i++ {
		for _, harf := range []rune(a[i]) {
			z01.PrintRune(harf)
		}
		z01.PrintRune('\n')
	}
}
