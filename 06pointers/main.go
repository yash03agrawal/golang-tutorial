package main

import "fmt"

func main() {

	num := 23

	var ptr = &num

	fmt.Println("ptr value:", ptr)
	fmt.Println("*ptr value", *ptr)
	fmt.Println("&ptr value", &ptr)

	*ptr = *ptr + 2
	fmt.Println("num value", num)
}
