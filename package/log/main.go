package main

import (
	"bytes"
	"crypto/x509/pkix"
	"log"
	"os"
	"path/filepath"
)

func example1() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	file, _ := os.OpenFile("log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)

	logger.SetOutput(file)
	logger.Println("dd")
	defer file.Close()
}

func example2() {
	rootdir, _ := os.Getwd()
	ofile := filepath.Join(rootdir, "/lijie.io/go_study_code/package/log/fff.log")
	file, _ := os.OpenFile(ofile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)
	log.Println("Hello world!")
	defer file.Close()
}

//
func main() {
	example1()
	example2()
	pkix.AlgorithmIdentifier{}
}
