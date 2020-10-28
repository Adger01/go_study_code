package main

import "fmt"

func funPrintln() {
	fmt.Println("Println")

	n, err := fmt.Println("李杰")
	fmt.Println(n)
	fmt.Println(err)
	fmt.Println("-----------------")
}

func funSprintf() {
	//Sprintf根据格式说明符格式化并返回结果字符串。
	fmt.Println("Sprintf")

	var i int
	s := fmt.Sprintf("i type : %T", i)
	fmt.Print(s)
	fmt.Println("")
	fmt.Println("-----------------")
}

func funPrintf() {

	i := 3
	fmt.Printf("整型：按值的本来值输出：%v \r\n",i)
	fmt.Printf("整型：结构体字段名和值进行展开：%+v \r\n",i)
	fmt.Printf("整型：输出 Go 语言语法格式的值：%#v \r\n",i)
	fmt.Printf("整型：类型：%T \r\n",i)//
	fmt.Printf("整型：二进制：%b \r\n",i)//
	fmt.Printf("整型：八进制：%o \r\n",i)//
	fmt.Printf("整型：十六进制：%x \r\n",i)//



	//数组
	var j [3]int
	j[0] = 4
	j[1] = 2
	j[2] = 5
	fmt.Printf("数组：按值的本来值输出：%v \r\n",j)// [4 2 5]
	fmt.Printf("数组：结构体字段名和值进行展开：%+v \r\n",j)//[4 2 5]
	fmt.Printf("数组：输出 Go 语言语法格式的值：%#v \r\n",j)//[3]int{4, 2, 5}
	fmt.Printf("数组：输出 Go 语言语法格式的类型和值：%T \r\n",j)//[3]int
	fmt.Printf("数组：二进制:%b \r\n",j)//[100 10 101]


	//切片
	k := []int{4,3,2}
	fmt.Printf("切片：按值的本来值输出: %v \r\n",k)
	fmt.Printf("切片：结构体字段名和值进行展开: %+v \r\n",k)
	fmt.Printf("切片：输出 Go 语言语法格式的值: %#v \r\n",k)
	fmt.Printf("切片：输出 Go 语言语法格式的类型和值: %T \r\n",k)
	fmt.Printf("切片：二进制: %b \r\n",k)
	//fmt.Println("")
	//fmt.Println("-----------------")


	//map
	l := map[string]string{
		"name":"jie",
		"age":"40",
		"gender":"male",
	}
	fmt.Printf("map：按值的本来值输出: %v \r\n",l)
	fmt.Printf("map：结构体字段名和值进行展开: %+v \r\n",l)
	fmt.Printf("map：输出 Go 语言语法格式的值: %#v \r\n",l)
	fmt.Printf("map：输出 Go 语言语法格式的类型和值: %T \r\n",l)
	fmt.Printf("map：二进制: %b \r\n",l)

	//struct
	type St struct {
		age int
		name string
		score map[string]int
	}

	var s St
	s.age = 2
	s.name = "lijie"
	s.score = map[string]int{
		"yuwen":99,
		"shuxue":45,
	}
	fmt.Printf("struct：按值的本来值输出: %v \r\n",s)
	fmt.Printf("struct：结构体字段名和值进行展开: %+v \r\n",s)
	fmt.Printf("struct：输出 Go 语言语法格式的值: %#v \r\n",s)
	fmt.Printf("struct：输出 Go 语言语法格式的类型和值: %T \r\n",s)
	fmt.Printf("struct：二进制: %b \r\n",s)



}

func main() {
	funPrintln()

	funSprintf()

	funPrintf()

		var (
			name    string
			age     int
			married bool
		)
		fmt.Scanln(&name, &age, &married)
		fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
