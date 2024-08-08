package main

import (
	"os"

	"github.com/01-edu/z01"
)

func IsNotdup(sja rune, result string) bool {
	for _, word := range result {
		if sja == word {
			return false
		}
	}
	return true
}

func mainuuuu() {
	if len(os.Args) != 3 {
		return
	}
	a := os.Args[1]
	b := os.Args[2]
	var change string
	for _, okk := range a {
		if IsNotdup(okk, change) {
			for _, okkk := range b {
				if okk == okkk {
					change += string(okk)
					break
				}
			}
		}
	}
	for _, looo := range change {
		z01.PrintRune(looo)
	}
	z01.PrintRune('\n')
}
