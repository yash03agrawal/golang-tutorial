package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	dice := rand.Intn(6) + 1

	switch dice {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	case 3:
		fmt.Println("case 3")
		fallthrough
	case 4:
		fmt.Println("case 4")
		fallthrough
	case 5:
		fmt.Println("case 5")
	case 6:
		fmt.Println("case 6")
	default:
		fmt.Println("not sure")
	}
}
