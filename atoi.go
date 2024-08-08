package main

import "fmt"

func Atoi(str string) int {
	result := 0
	negative := false
	for _, word := range str {
		if word == '-' {
			negative = true
			continue
		}
		if word < '0' || word > '9' {
			return 0
		}
		result = result*10 + int(word-'0')
	}
	if negative {
		return -result
	}
	return result
}

func mainiiio() {
	// Test cases
	fmt.Println(Atoi("123"))  // Output: 123, <nil>
	fmt.Println(Atoi("-456")) // Output: -456, <nil>
	fmt.Println(Atoi("789"))  // Output: 789, <nil>
	fmt.Println(Atoi("12a3")) // Output: 0, invalid character found
}
