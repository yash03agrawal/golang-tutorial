package main

import "fmt"

// defer follows Stack - LIFO

func main() {

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("World")
	fmt.Println("Hello")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
