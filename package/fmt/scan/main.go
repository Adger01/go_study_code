package main

import(
	"fmt"
)

func main() {
	type Student struct {
		name string
		score int
	}

	var s Student
	fmt.Println("请输入姓名：")
	fmt.Scan(&s.name)

	fmt.Println("请输入年龄：")
	fmt.Scan(&s.score)

	fmt.Println(s)
}

//Scan和Scanln基本相同，唯一区别是当读取多个变量当时候，遇到换行符Scanln会直接结束，未读到输入值的变量为零值；Scan会等待，直到输入的值满足参数的个数后再遇到换行符才会结束。