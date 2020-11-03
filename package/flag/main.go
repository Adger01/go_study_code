package main

import (
	"flag"
	"fmt"
)

var (
	intflag    int
	boolflag   bool
	stringflag string
)

func init() {
	flag.IntVar(&intflag, "intflag", 0, "intflag")
	flag.BoolVar(&boolflag, "boolflag", false, "boolflag")
	flag.StringVar(&stringflag, "stringflag", "", "stringflag")
}

func main() {
	flag.Parse()

	fmt.Println(intflag)
	fmt.Println(boolflag)
	fmt.Println(stringflag)

}
