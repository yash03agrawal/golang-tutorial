package main

import "fmt"

func main() {

	days := []string{"monday", "tuesday", "wednesday"}

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, val := range days {
		fmt.Printf("index: %v and value: %v\n", index, val)
	}

	// break, continue, goto are used in the same way
}
