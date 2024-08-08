package main

import (
	"fmt"
)

// FromTo takes two integers and returns a string showing the range of numbers from the first to the second
func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	var result string
	if from <= to {
		for i := from; i <= to; i++ {
			if i < 10 {
				result += "0" + Itooa(i)
			} else {
				result += Itooa(i)
			}
			if i < to {
				result += ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {
			if i < 10 {
				result += "0" + Itooa(i)
			} else {
				result += Itooa(i)
			}
			if i > to {
				result += ", "
			}
		}
	}

	return result + "\n"
}
func Itooa(n int) string {
	if n == 0 {
		return "0"
	}
	negative := false
	var digit int
	var digits string
	if n < 0 {
		negative = true
		n = -n
	}
	if n > 0 {
		digit = n % 10
		digits = string(rune(digit)+'0') + digits
		n /= 10
	}
	if negative {
		digits = "-" + digits
	}
	return digits
}

func maiooon() {
	// fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, -9))
}
