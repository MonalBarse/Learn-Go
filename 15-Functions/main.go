package main

import "fmt"

func main() { // Main function is the entry point of the program
	// We don't need to call the main function. It's called automatically when we run the program
	fmt.Println("Functions in Go")
	greetings()

	result1 := add(10, 20)
	fmt.Printf("Result of addition is %v and the type is %T\n", result1, result1)

	result2 := divide(10, 3)
	fmt.Printf("Result of division is %v and the type is %T\n", result2, result2)

	result3 := superAdd(10.01, 20.05, 30.3, 40.3, 50.5)
	fmt.Printf("Result of superAdd is %v and the type is %T\n", result3, result3)
}

// --------------------------------------------------- //

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

// ---------There's more about functions which we will see in future modules-------- //
