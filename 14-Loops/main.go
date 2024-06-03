package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in Go")
	fmt.Println("Go has only one looping construct, the for loop.")

	// For loop has serverl variations
	// 1. for loop with single condition
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("=====================================")
	// 2. for loop with multiple conditions
	fmt.Println("i j")
	for i, j := 0, 100; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	fmt.Println("=====================================")
	// 3. for loop as while loop (no while loop in Go)
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// 4. for loop as infinite loop
	// for {
	//   fmt.Println("Infinite loop")
	// }

	// 5. for loop with range
	// range keyword is used to iterate over elements in a variety of data structures like arrays, slices, maps, strings, etc.
	// It returns two values, the index and the value of the element at that index.
	// It's similar to JS for..of loop or Java for-each loop

	fmt.Println("=====================================")
	mySlice := []int{55, 20, 83, 94, 665}
	fmt.Println("index value")
	for index, value := range mySlice {
		fmt.Println(index, value)
	}

	fmt.Println("=====================================")

	myNum := 10

	for myNum > 0 {
		if myNum == 2 {
			goto lco
		}
		if myNum == 5 {
			fmt.Println("Skipping this iteration with continue keyword")
			myNum--
			continue
		}
		fmt.Println("Vlaue is:", myNum)
		myNum--
	}
lco:
	fmt.Println("LCO")
}
