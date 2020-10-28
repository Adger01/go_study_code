package main

import (
	"fmt"
	"os"
)

//第一种方式，使用os和for循环
func method1() {
	l := len(os.Args)
	for i := 1; i < l; i++ {
		fmt.Println(os.Args[i])
	}
}

//第二种方式，使用range和os
func method2() {
	var s, step string
	s = ""
	step = ""
	for _, v := range os.Args[1:] {
		s += step + v
		step = " "
	}
	fmt.Println(s)
}

func main() {
	method1()
	//method2()
}
