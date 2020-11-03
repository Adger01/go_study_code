package main

import (
	"log"
)

/**
在使用go开发项目的时候，经常要用到log包的输出格式，go自带了默认的输出格式是时间+内容，同时也提供了自定义格式。
其中log包含以下几个输出选项
const (
	Ldate         = 1 << iota     // 日期:  2009/01/23
	Ltime                         // 时间:  01:23:23
	Lmicroseconds                 // 微秒:  01:23:23.123123.
	Llongfile                     // 路径+文件名+行号: /a/b/c/d.go:23
	Lshortfile                    // 文件名+行号:   d.go:23
	LUTC                          // 使用标准的UTC时间格式
	Lmsgprefix                    // 将“前缀”从行首移到消息之前
	LstdFlags     = Ldate | Ltime // 默认
)
*/
func main() {
	LstdFlags()
	Lmsgprefix()
	LUTC()
	Lshortfile()
	Llongfile()
	Lmicroseconds()
	Ltime()
	Ldate()
	LdateAndLtime()
}

func LstdFlags() {
	log.Println("这是默认的格式\n")
	//2020/10/29 14:49:08 这是默认的格式
}

func Lmsgprefix() {
	log.SetFlags(log.Lmsgprefix)
	log.Println("这是Lmsgprefix的格式\n")
	//这是Lmsgprefix的格式
}

func LUTC() {
	log.SetFlags(log.LUTC)
	log.Println("这是LUTC的格式\n")
	//这是LUTC的格式
}

func Lshortfile() {
	log.SetFlags(log.Lshortfile)
	log.Println("这是Lshortfile的格式\n")
	//main.go:47: 这是Lshortfile的格式
}

func Llongfile() {
	log.SetFlags(log.Llongfile)
	log.Println("这是Llongfile的格式")
	///Users/xxx/go/src/xxx.io/go_study_code/package/log/flag/main.go:54: 这是Llongfile的格式
}

func Lmicroseconds() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("这是Lmicroseconds的格式")
	//14:53:54.240599 这是Lmicroseconds的格式
}

func Ltime() {
	log.SetFlags(log.Ltime)
	log.Println("这是Ltime的格式")
	//14:55:07 这是Ltime的格式
}

func Ldate() {
	log.SetFlags(log.Ldate)
	log.Println("这是Ldate的格式")
	//2020/10/29 这是Ldate的格式
}

func LdateAndLtime() {
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("这是Ldate+Ltime的格式")
	//2020/10/29 14:56:30 这是Ldate+Ltime的格式
}
