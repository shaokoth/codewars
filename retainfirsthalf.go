package main

import (
	"fmt"
)

func RetainFirstHalf(str string) string {
	for i := 0; i <= len(str)-1; i++ {
		if len(str) == 0 {
			return ""
		} else if len(str) == 1 {
			return str
		} else {
			return str[:len(str)/2]
		}
	}
	return ""
}

func mainhjg() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
