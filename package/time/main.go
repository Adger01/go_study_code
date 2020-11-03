package main

import (
	"fmt"
	"time"
)

func main() {
	//打印当前时间
	t := time.Now()

	//年
	fmt.Println(t.Year())

	//月
	fmt.Println(t.Month())

	//日
	fmt.Println(t.Day())

	//年月日
	fmt.Println(t.Date())

	//时
	fmt.Println(t.Hour())

	//分
	fmt.Println(t.Minute())

	//秒
	fmt.Println(t.Second())

	//时间戳
	fmt.Println(t.Unix())

	//一年中第几天
	fmt.Println(t.YearDay())

	//星期几 英文
	fmt.Println(t.Weekday())

	//获取一年中第几周
	fmt.Println(t.ISOWeek())

	timer := time.Timer{}
	fmt.Println(timer.Stop())

}
