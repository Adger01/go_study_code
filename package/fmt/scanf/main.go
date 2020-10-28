package main

import "fmt"

func main(){
	type Student struct {
		name string
		score int
	}
	var s Student
	//fmt.Scanf必须按照第一个参数的format个数输入

	fmt.Scanf("姓名：%s",&s.name)
	fmt.Scanf("年龄：%d",&s.score)

	fmt.Println(s)
}
