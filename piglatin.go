package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Isvowel(word rune) bool {
	vow := "aeiouAEIOU"
	for _, char := range vow {
		if word == char {
			return true
		}
	}
	return false
}

func Piglatin(s string) string {
	if Isvowel(rune(s[0])) {
		return s + "ay"
	}
	for i, six := range s {
		if Isvowel(six) {
			return s[i:] + s[:i] + "ay"
		}
	}
	return "No vowels"
}

func mainlll() {
	if len(os.Args) != 2 {
		return
	}
	result := Piglatin(os.Args[1])

	for _, jaa := range result {
		z01.PrintRune(jaa)
	}
	z01.PrintRune('\n')
}
