package main

import "fmt"

const LoginToken string = "lkdfkljdf" // public variable because of first letter as capital letter

func main() {
	var username string = "yash"
	fmt.Println(username)
	fmt.Printf("Type of username is %T\n", username)

	var isUserLoggedIn bool = true
	fmt.Println(isUserLoggedIn)
	fmt.Printf("Type of isUserLoggedIn is %T\n", isUserLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Type of smallVal is %T\n", smallVal)

	var floatVal float32 = 255.333343434
	fmt.Println(floatVal)
	fmt.Printf("Type of floatVal is %T\n", floatVal)

	// default value
	var intNum int
	fmt.Println(intNum)
	fmt.Printf("Type of intNum is %T\n", intNum)

	// implicit type
	var numb = 3
	fmt.Println(numb)

	// no var style - only allowed inside methods not as global
	numberOfUsers := 100
	fmt.Println(numberOfUsers)
	fmt.Printf("Type of numberOfUsers is %T\n", numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Type of LoginToken is %T\n", LoginToken)

}
