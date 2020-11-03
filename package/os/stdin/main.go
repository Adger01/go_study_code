package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func example(file string) {
	counts := make(map[string]int)

	//打开文件
	fileInfo, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//扫描文件内容
	sc := bufio.NewScanner(fileInfo)

	for sc.Scan() {
		counts[sc.Text()]++
	}

	fmt.Println(counts)
}

func example2(file string) {
	fileInfo, err := os.Open(file)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	r := bufio.NewReader(fileInfo)

	var counts map[string]int
	counts = make(map[string]int)
	for {
		b, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err:", err.Error())
			return
		}
		str := strings.Replace(string(b), "\n", "", -1)

		counts[str]++
	}

	fmt.Printf("%#v", counts)
}

//题目：读取一个文件，输出重复的字符串以及行数，放到一个map中
func main() {
	file := "./a.txt"

	example(file)

	example2(file)
}
