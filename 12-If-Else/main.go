package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("IF-ELSE in Go!")
	fmt.Println("-------------------------")
	fmt.Println("The loginCount is 10")
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
	fmt.Println("To show that we can also initialize a variable in the if statement")
	// Check if the user input integer is even or odd

	fmt.Println("Enter a number to check if it is even or odd")
	if num, _ := ReadInt64(); num == 0 { // ReadInt64() is from the helper function below
		fmt.Printf("%d is neither even nor odd\n", num) // %d is a placeholder for the integer stands for decimal
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
		return 0, errors.New("error reading input") // we return 0 cause we need to return an int64
	}

	input = strings.TrimSpace(input)

	num, err := strconv.ParseInt(input, 10, 64) // 10 is the base and 64 is the bit size (ParseInt takes 3 arguments - input, base, bit size)
	// Eg: ParseInt("111", 2, 64) will return 7
	if err != nil {
		return 0, fmt.Errorf("error converting input to int64: %v", err)
	}

	return num, nil // Return nil if no error (the other errors have been handled by fmt.Errorf)
}
