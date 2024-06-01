package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	fmt.Println(result) // Exactly 10 login count

	fmt.Println("-------------------------")

	// Check if the user input integer is even or odd

	if num, _ := ReadInt64(); num == 0 {
		fmt.Printf("%d is neither even nor odd\n", num)
	} else if num%2 == 0 {
		fmt.Printf("%d is an even number\n", num)
	} else {
		fmt.Printf("%d is an odd number\n", num)
	}
	fmt.Println("-------------------------")
}

func ReadInt64() (int64, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("error reading input: %v", err)
	}

	input = strings.TrimSpace(input)

	num, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting input to int64: %v", err)
	}

	return num, nil // Return nil if no error (the other errors have been handled by fmt.Errorf)
}
