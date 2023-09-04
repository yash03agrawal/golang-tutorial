package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"mango"}
	fruitList = append(fruitList, "banana", "grapes")
	fmt.Println("fruit list:", fruitList)

	fruitList = append(fruitList[:1])
	fmt.Println("fruit list:", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 321
	highScores[2] = 456
	highScores[3] = 897

	// error
	// highScores[4] = 322

	highScores = append(highScores, 121, 321)
	fmt.Println("high scores", highScores)

	sort.Ints(highScores)
	fmt.Println("high scores", highScores)
	fmt.Println("is sorted", sort.IntsAreSorted(highScores))

	// removing values
	var lang = []string{"python", "c", "rust", "ruby"}
	fmt.Println("languages", lang)

	var index int = 2
	lang = append(lang[:index], lang[index+1:]...)
	fmt.Println("languages", lang)
}
