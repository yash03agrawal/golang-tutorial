package main

import "fmt"

func main() {
	greeter()

	res := adder(3, 5)
	fmt.Println("addition is:", res)

	result, _ := proAdder(1, 2, 3, 4, 5, 6)
	fmt.Println("result:", result)

}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hello from pro adder"
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func greeter() {
	fmt.Println("Hello world!")
}
