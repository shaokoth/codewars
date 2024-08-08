package main

import "fmt"

func PrintMemory(arr [10]byte) {
	for i, r := range arr {
		var s string
		if r == 0 {
			s = "00"
		}
		for r > 0 {
			if r%16 == 0 {
				s = "0" + s
			}
			if r%16 == 1 {
				s = "1" + s
			}
			if r%16 == 2 {
				s = "2" + s
			}
			if r%16 == 3 {
				s = "3" + s
			}
			if r%16 == 4 {
				s = "4" + s
			}
			if r%16 == 5 {
				s = "5" + s
			}
			if r%16 == 6 {
				s = "6" + s
			}
			if r%16 == 7 {
				s = "7" + s
			}
			if r%16 == 8 {
				s = "8" + s
			}
			if r%16 == 9 {
				s = "9" + s
			}
			if r%16 == 10 {
				s = "a" + s
			}
			if r%16 == 11 {
				s = "b" + s
			}
			if r%16 == 12 {
				s = "c" + s
			}
			if r%16 == 13 {
				s = "d" + s
			}
			if r%16 == 14 {
				s = "e" + s
			}
			if r%16 == 15 {
				s = "f" + s
			}
			r /= 16
		}
		fmt.Print(s)
		if i == 3 || i == 7 || i == 9 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
	for _, c := range arr {
		if c < 32 || c > 126 {
			fmt.Print(string('.'))
		}
		fmt.Print(string(c))
	}
	fmt.Println()
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
