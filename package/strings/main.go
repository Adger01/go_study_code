package main

import (
	"fmt"
	"strings"
)

//strings.json

func StrPlus1(a []string) string {
	var s, sep string
	for i := 0; i < len(a); i++ {
		s += sep + a[i]
		sep = " "
	}
	return s
}

func StrPlus2(a []string) string {
	return strings.Join(a," ")
}
func main() {
	a := []string{
		"li",
		"jie",
	}

	fmt.Println(a)
	fmt.Println(StrPlus1(a))
	fmt.Println(StrPlus2(a))
}