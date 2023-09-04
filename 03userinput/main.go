package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	// https://pkg.go.dev/
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating: ")

	// comma ok || comma err
	input, _ := reader.ReadString('\n')
	fmt.Println("Rating given is:", input)
	fmt.Printf("Type of rating is: %T", input)
}
