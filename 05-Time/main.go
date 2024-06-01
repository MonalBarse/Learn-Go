package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("We will study the time package in Go.")

	presentTime := time.Now()
	fmt.Println("The current time is: ", presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2023, time.August, 15, 10, 0, 0, 0, time.UTC)
	fmt.Println("The date is: ", createdDate.Format("01-02-2006 Monaday"))
	// WE will run go env to check our OS and build for the Windows OS and Mac OS by running `GOOS=windows go build`
	// and `GOOS=darwin go build` respectively
}
