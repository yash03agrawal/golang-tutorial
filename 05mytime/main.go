package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time")

	presentTime := time.Now()
	fmt.Println("Present time:", presentTime)
	fmt.Println("Formatted present time:", presentTime.Format("02-01-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 03, 23, 00, 00, 00, time.UTC)
	fmt.Println("created date", createdDate)
	fmt.Println("formatted created date", createdDate.Format("02-01-2006 Monday"))

	/**
	 * to build for different os:
	 * mac (self system): go build
	 * windows: GOOS="windows" go build
	 * linux: GOOS="linux" go build
	 */
}
