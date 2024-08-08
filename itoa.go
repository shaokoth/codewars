package main

import "fmt"

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}
	var result string
	for n > 0 {
		digit := n%10
		result = string(rune(digit) + '0') + result
		n /= 10
	}
	if negative {
		result = string(rune('-')) + result
	}
	return result
}

func maiddsn() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
