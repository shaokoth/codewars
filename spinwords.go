package main

import (
	"fmt"
	"strings"
)

func Spinwords(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		if len(word) >= 5 {
			var reversedWord string
			for j := len(word) - 1; j >= 0; j-- {
				reversedWord += string(word[j])
			}
			words[i] = reversedWord
		}
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(Spinwords("This is a khhhkkhhtest"))
	fmt.Println(Spinwords("This is another test"))
}
