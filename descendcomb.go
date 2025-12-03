package piscine

import (
	"fmt"
)

func DescendComb() {
	for a := 99; a >= 0; a-- {
		for b := a - 1; b >= 0; b-- {
			showNumber(a)
			fmt.Print(' ')
			showNumber(b)
			if a != 1 || b != 0 {
				fmt.Print(',')
				fmt.Print(' ')
			}
		}
	}
}

func showNumber(n int) {
	first := n / 10
	second := n % 10
	// print digits as characters (same effect as z01.PrintRune)
	fmt.Print(string(rune(first + 48)))
	fmt.Print(string(rune(second + 48)))
}
