package main

import "fmt"

func main() {

	john := User{"John", 16, true}
	fmt.Println("John details", john)
	fmt.Printf("john details %+v\n", john)
	fmt.Printf("john name: %v\n", john.Name)
}

type User struct {
	Name       string
	Age        int
	IsVerified bool
}
