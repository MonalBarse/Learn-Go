package main

import "fmt"

func greetings() {
	fmt.Println("Hello from greetings function")
}

// --------------------------------------------------- //

func add(a int, b int) float64 {
	x := float64(a) + float64(b) // Convert a and b to float64 before addition
	fmt.Printf("type of x is %T \n", x)

	return x
}

// --------------------------------------------------- //

func divide(a int, b int) float64 {
	x := float64(a) / float64(b) // Convert a and b to float64 before division
	fmt.Printf("type of x is %T \n", x)

	return x
}

// --------------------------------------------------- //
func superAdd(val ...float64) float64 {
	fmt.Printf("type of val is %T \n", val)

	result := 0.0
	for _, v := range val {
		result += v
	}

	return result
}
func reverse(i ...int64) []int64 {
	// Example: myDefer(1, 2, 3, 4) => [4, 3, 2, 1]

	// first we create a slice to store the reversed integers
	reversed := make([]int64, 0)
	// iterate
	for _, v := range i {
		// prepend the integer to the slice
		reversed = append([]int64{v}, reversed...)
	}
	return reversed
}

// ---------There's more about functions which we will see in future modules-------- //
func main() { // We don't need to call the main function. It's called automatically when we run the program
	fmt.Println("Functions in Go")
	fmt.Println("Main function is the entry point of the program")

	fmt.Println("----------------------------")
	greetings()

	fmt.Println("----------------------------")
	result1 := add(10, 20)
	fmt.Printf("Result of addition is %v and the type is %T\n", result1, result1)

	fmt.Println("----------------------------")
	result2 := divide(10, 3)
	fmt.Printf("Result of division is %v and the type is %T\n", result2, result2)

	fmt.Println("----------------------------")
	result3 := superAdd(10.01, 20.05, 30.3, 40.3, 50.5)
	fmt.Printf("Result of superAdd is %v and the type is %T\n", result3, result3)

	fmt.Println("----------------------------")
	result4 := reverse(99, 88, 77, 66, 55)
	fmt.Printf("Result of reverse is %v and the type is %T\n", result4, result4)
}

// --------------------------------------------------- //
