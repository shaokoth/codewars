package main

import "fmt"

func CanJump(steps []uint) bool {
	if len(steps) == 0 {
		return false
	}
	if len(steps) == 1 {
		return true
	}
	n := len(steps)
	currentPosition := 0
	for currentPosition < n-1 {
		if steps[currentPosition] == 0 {
			return false
		}
		currentPosition += int(steps[currentPosition])
		if currentPosition == n-1 {
			return true
		} else if currentPosition >= n {
			return false
		}
	}
	return false
}

func maidggn() {
	input1 := []uint{2, 3, 1, 1, 4, 1}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4, 5}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
