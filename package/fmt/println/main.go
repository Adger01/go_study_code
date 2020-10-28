package main

import "fmt"

func main() {
	t := 3
	tp := &t
	fmt.Printf("t地址是：%d，类型%T \n",&t,&t)
	fmt.Printf("类型是：%T，值是%d\n",*tp,*tp)
}


