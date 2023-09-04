package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "Content of file"
	fileName := "./sample.txt"
	file, err := os.Create(fileName)

	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length written is:", length)

	// use fo defer
	defer file.Close()

	readFile(fileName)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(file string) {
	dataBytes, err := ioutil.ReadFile(file)
	checkNilErr(err)

	fmt.Println("file content:", string(dataBytes))
}
