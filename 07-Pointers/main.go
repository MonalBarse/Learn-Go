package main

import "fmt"

func main() {
	fmt.Println("File for Pointers in Go.")
	// We will study the pointers in Go.
	// A pointer is a variable that stores the memory address of another variable.
	// Creating a pointer
	var ptr *int
	// The * operator is used to declare a pointer.

	fmt.Println("The value of ptr is: ", ptr) // The value of ptr is: <nil>

	myNumber := 55
	var ptr1 = &myNumber                          // The & operator is used to get the address of a variable.
	fmt.Println("The value of ptr1 is: ", ptr1)   // The value of ptr1 is: 0xc0000b4008
	fmt.Println("The value of *ptr1 is: ", *ptr1) // The value of *ptr1 is: 55
	// Pointers are the only thing in GO that is passed by reference. Everything else is passed by value. eg. all the primitve data types and arrays structs etc are passed by value. But slices can behave like pass by reference as they use pointers internally.
}
