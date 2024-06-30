package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in Go")
	fmt.Println("Go has only one looping construct, the for loop.")

	// For loop has serverl variations
	// 1. for loop with single condition
	fmt.Println("======= for with single condition ==========")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("======= for with multiple conditions ==========")
	// 2. for loop with multiple conditions
	fmt.Println("i j")
	for i, j := 0, 100; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	fmt.Println("========= for loop as while loop ==========")
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

	fmt.Println("========== for loop with range ==========")
	mySlice := []int{55, 20, 83, 94, 665}
	fmt.Println("index value")
	for index, value := range mySlice {
		fmt.Println(index, value)
	}

	fmt.Println("========== for loop with range on map ==========")

	/* myNum := 10

		for myNum > 0 {
			if myNum == 2 {
				goto lco // This is not recommended to use goto keyword; it makes the code hard to read and understand instead use break, continue or return
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
		fmt.Println("Jumped to lco") // This is the label where the goto keyword jumps to. This is of data type label and is used with goto keyword and no where else.
	*/
	// Instead of using goto keyword we can use break, continue or return keyword to achieve the same result.
	// break keyword is used to break out of the loop
	// continue keyword is used to skip the current iteration and move to the next iteration
	// return keyword is used to return from the function

	myNum2 := 10

	for myNum2 > 0 {
		if myNum2 == 2 {
			break
		}
		if myNum2 == 5 {
			fmt.Println("Skipping this iteration with continue keyword")
			myNum2--
			continue
		}
		fmt.Println("Vlaue is:", myNum2)
		myNum2--
	}
	fmt.Println("The loop is broken using break keyword")

}
