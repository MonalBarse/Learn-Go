package main

import (
	"fmt"
)

func main() {

	fmt.Println("IF-ELSE in Go!")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out!"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)
}
