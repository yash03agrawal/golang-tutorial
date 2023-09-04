package main

import "fmt"

func main() {
	count := 10
	var result string

	if count > 10 {
		result = "greater"
	} else if count < 10 {
		result = "less"
	} else {
		result = "equal"
	}

	fmt.Println(result)

	if num := 10; num == 10 {
		fmt.Println("checked")
	}
}
