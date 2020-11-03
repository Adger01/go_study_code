package main

import (
	"fmt"
	"os"
)

func example(file string) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("name:", fileInfo.Name())
	fmt.Println("isdir:", fileInfo.IsDir())
	fmt.Println("mode:", fileInfo.Mode())
	fmt.Println("modTime:", fileInfo.ModTime())
	fmt.Println("size:", fileInfo.Size())
	fmt.Println("sys:", fileInfo.Sys())
}

func main() {
	file := "./a.txt"
	example(file)
}
