package main

import "fmt"

func main() {

	fmt.Println("----------- Defer in Golang -----------")
	// Defer is used to delay the execution of a statement until the end of the function
	// It's used to ensure that a function call is performed later in a program's execution

	// Defer is often used when resources need to be freed in some way

	defer fmt.Println("Last function")
	defer fmt.Println("Last second function")
	first()
	second()
	myDefer()
	/*
	   Output:
	*/
}

// --------------------------------------------------- //
func first() {
	fmt.Println("First function")
}

func second() {
	defer fmt.Println("Second function part 2")
	defer fmt.Println("Second function part 1")

	fmt.Println("Second function")
}

// --------------------------------------------------- //

func myDefer() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}

// ----------- Defer in Golang -----------
// First function
// Second function
// Second function part 1
// Second function part 2
// 5
// 4
// 3
// 2
// 1
// Last second function
// Last function
