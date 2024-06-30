package main

import "fmt"

func main() {
	fmt.Println("Understanding Pointers in Go")

	// 1. Declaration and Initialization
	var ptr *int // Declaration without initialization (nil pointer)
	fmt.Println("1. Declared ptr without initialization:", ptr)

	myNumber := 10
	ptr = &myNumber // Initialize ptr with the address of myNumber
	fmt.Println("Initialized ptr with address of myNumber:", ptr)
	fmt.Println("Value stored in ptr:", *ptr)

	// 2. Dereferencing
	fmt.Println("\n2. Dereferencing ptr to get the value of myNumber:", *ptr)

	// 3. Passing Pointers to Functions
	fmt.Println("\n3. Passing Pointers to Functions")
	fmt.Println("Before function call, myNumber:", myNumber)
	increment(&myNumber) // Pass the address of myNumber to increment function
	fmt.Println("After function call, myNumber:", myNumber)

	// 4. Pointers with Structs
	fmt.Println("\n4. Pointers with Structs")
	p := Person{Name: "Alice", Age: 25}
	ptrToPerson := &p // Pointer to struct
	fmt.Println("Initial struct value:", p)
	fmt.Println("Pointer to struct:", ptrToPerson)
	fmt.Println("Accessing struct fields through pointer:", (*ptrToPerson).Name, ptrToPerson.Age)

	// 5. Nil Pointers and Error Handling
	fmt.Println("\n5. Nil Pointers and Error Handling")
	var nilPtr *int // nil pointer
	if nilPtr == nil {
		fmt.Println("nilPtr is nil")
	} else {
		fmt.Println("Value stored in nilPtr:", *nilPtr)
	}

	// 6. Pointers vs. Value Types
	fmt.Println("\n6. Pointers vs. Value Types")
	val := 100
	ptrVal := &val // Pointer to value
	fmt.Println("Value of val:", val)
	fmt.Println("Address of val:", ptrVal)
	fmt.Println("Value accessed through pointer:", *ptrVal)
	*ptrVal = 200
	fmt.Println("Updated value of val through pointer:", val)

	// Example with different types
	fmt.Println("\nExample with different types:")
	var str string = "Hello"
	strPtr := &str
	fmt.Println("String value:", str)
	fmt.Println("String address:", strPtr)
	fmt.Println("String value through pointer:", *strPtr)
	*strPtr = "World" // Change value through pointer
	fmt.Println("Updated string value through pointer:", str)

	// 7. Using Pointers with Arrays
	fmt.Println("\n7. Using Pointers with Arrays")
	arr := [3]int{1, 2, 3}
	arrPtr := &arr
	fmt.Println("Array values:", arr)
	fmt.Println("Array pointer address:", arrPtr)
	// Access array elements through pointer
	fmt.Println("First element via pointer:", (*arrPtr)[0])
	fmt.Println("Second element via pointer:", (*arrPtr)[1])
	fmt.Println("Third element via pointer:", (*arrPtr)[2])

	// 8. Understanding Memory Address and Pointer Arithmetic
	fmt.Println("\n8. Understanding Memory Address and Pointer Arithmetic")
	a := 10
	b := 20
	ptrA := &a
	ptrB := &b
	fmt.Println("Address of a:", ptrA)
	fmt.Println("Address of b:", ptrB)

	// 9. Garbage Collection and Memory Management
	fmt.Println("\n9. Garbage Collection and Memory Management")
	// Go's garbage collector automatically manages memory.
	// Explicit memory management is generally not required in Go.
	// Example: Creating and deleting large struct to demonstrate GC.
	type LargeStruct struct {
		data [1e6]byte
	}
	// large := LargeStruct{}
	fmt.Println("Large struct created")
	// No need to free memory explicitly. GC will handle it.

	fmt.Println("\nEnd of Pointers Program")
}

// Struct definition for example with structs
type Person struct {
	Name string
	Age  int
}

func increment(ptr *int) {
	*ptr++ // Increment the value at the pointer's address
}
