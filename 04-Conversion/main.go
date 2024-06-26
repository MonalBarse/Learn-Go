package main

import (
	"bufio"   // for reading input specifically from the console
	"fmt"     // for printing the output to the console
	"os"      // the bufio package uses the os package to read the input from the console bufio for file i/o is bufio.NewReader(file)
	"strconv" // for converting the string to a float64
	"strings" // strings package for manipulating strings like trimming the white spaces from the input text
)

func main() {
	fmt.Println("Welcome to our Shop!")
	fmt.Println("Rate your last purchase: ")

	reader := bufio.NewReader(os.Stdin) // THis line will read the input from the user
	input, _ := reader.ReadString('\n') // This line will read the input from the user until the user press enter
	fmt.Println("You have rated:", input)
	fmt.Println("Type of input:", fmt.Sprintf("%T", input)) // This line will print the type of the input which is a string. In below code we convert the string to a float64
	fmt.Println("-------------------")

	// numRating, err := strconv.ParseFloat(input, 64) // The strconv package connverts the string to a float64 or (if use ParseInt) to an int64 etc..
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // This line will remove the white spaces from the input
	if err != nil {
		fmt.Println("Please enter a valid number")
		return
	}
	fmt.Println("Final Rating:", numRating+1)

}
