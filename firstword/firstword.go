package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Arguments := os.Args

	if len(Arguments) != 2 {
		return
	}

	firstSpace := true
	word := ""
	for _, char := range Arguments[1] {
		if firstSpace && char == ' ' {
			continue
		}
		if char != ' ' {
			word += string(char)
			firstSpace = false
		} else {
			break
		}
	}
	PrintStr(word)
}

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
