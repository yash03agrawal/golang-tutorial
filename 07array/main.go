package main

import "fmt"

func main() {
	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "mango"
	fruitList[3] = "grapes"

	fmt.Println("fruit list:", fruitList)
	fmt.Println("length of fruits:", len(fruitList))

	var vegList = [3]string{"potato", "onion", "cabbage"}
	fmt.Println("veg list:", vegList)
	fmt.Println("length of vegList:", len(vegList))
}
