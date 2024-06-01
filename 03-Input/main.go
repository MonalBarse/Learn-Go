package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "UserInput"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number of elements in the array: ")

	// comma ok || err ok syntax
	/*
	   // This is the syntax for comma ok || err ok
	   input , _ := reader.ReadString('\n')
	   fmt.Println("Thanks for that,",input)
	   // Here we the input will be stored in input and if any error occurs it
	   // uses the _ to ignore the error ( or we can use err instead of _)
	*/

	input, err := reader.ReadString('\n') // Read the input string from the user till the user presses enter (i.e. '\n')
	if err != nil {
		fmt.Println("An error occurred while reading the input string.")
		return
	}
	fmt.Printf("The number you entered is %s, it is of type %T\n", input, input)
}
