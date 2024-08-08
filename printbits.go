package main

import (
	"fmt"
	"os"
	"strconv"
)

func mai5rn() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Print("00000000")
		return
	}

	if num < 0 || num > 255 {
		fmt.Print("00000000")
		return
	}

	fmt.Printf("%08b", num)
}
