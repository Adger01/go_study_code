package main

import (
	"fmt"
	"reflect"
)

type CustomStack []int

func PushStack(stk []int, v int) CustomStack {
	return append(stk, v)
}

func PopStack(stk []int) (CustomStack, int) {
	l := len(stk)
	if l == 0 {
		fmt.Println("列表为空")
		return stk, 0
	}
	return stk[:l-1], stk[l-1]
}
func main() {
	fmt.Println("start ...")

	st := make([]int, 0)
	fmt.Println(st)
	st = PushStack(st, 4)
	st, _ = PopStack(st)
	st = PushStack(st, 2)
	st = PushStack(st, 7)
	st = PushStack(st, 9)
	fmt.Println(st)
	st, _ = PopStack(st)
	fmt.Println(reflect.TypeOf(st))
	fmt.Println(st)
}
