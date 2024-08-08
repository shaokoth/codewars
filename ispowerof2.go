package main

import (
	"fmt"
	"os"
	"strconv"
)

func Ispowerof2(n int) bool{
return n > 0 && (n & (n-1)) == 0
}

func mafhfin() {
	if len(os.Args) != 2 {
		return
	}
	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("error converting to integer")
	}
	if Ispowerof2(number) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
