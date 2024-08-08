package main

import (
	"fmt"
	"os"
	"strings"
)

func Cleanstr(str string) string {
	strr := strings.TrimSpace(str)
	s := strings.Fields(strr)
	return strings.Join(s, " ")
}

func maddein() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	if args == "" {
		return
	}
	fmt.Println(Cleanstr(args))
}
