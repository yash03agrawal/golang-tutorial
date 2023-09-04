package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"

	fmt.Println("languages", languages)

	delete(languages, "PY")
	fmt.Println("languages", languages)

	//loops

	for key, value := range languages {
		fmt.Printf("key %v and value %v\n", key, value)
	}
}
