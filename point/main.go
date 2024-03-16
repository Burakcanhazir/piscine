package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	PrintString("x = ")
	PrintInt(points.x)
	PrintString(",")
	PrintString(" y = ")
	PrintInt(points.y)
	z01.PrintRune('\n')
}

func PrintInt(a int) {
	r := '0'
	if a/10 > 0 {
		PrintInt(a / 10)
	}
	for i := 0; i < a%10; i++ {
		r++
	}
	z01.PrintRune(r)
}

func PrintString(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
}
