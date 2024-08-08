package main

import (
	"fmt"
)

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
func PrevPrime(num int) int {
	for i := num; i >= 2; i-- {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}

func main44() {
	fmt.Println(PrevPrime(4))
	fmt.Println(PrevPrime(15))
}
