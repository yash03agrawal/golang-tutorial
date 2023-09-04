package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=233ksjdflkj"

func main() {

	fmt.Println("hello")

	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	query := result.Query()

	fmt.Printf("Type of query %T\n", query)

	fmt.Println(query["coursename"])

	for key, val := range query {
		fmt.Println("param is:", key, val)
	}

	// constructing url- always reference
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=john",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
