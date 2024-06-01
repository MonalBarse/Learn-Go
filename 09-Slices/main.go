package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Go")
	// Slices are a data structure in Go that provides a more powerful interface to sequences than arrays.
	// A slice is a lightweight structure that wraps and represents a portion of an array.
	// It is preferred over arrays because it is more flexible, convenient, and efficient.

	// Creating a slice Syntax : var sliceName []dataType
	var mySlice = []int{1, 2, 3, 4, 5}
	fmt.Printf("The type of mySlice is : %T \n", mySlice) // The type of mySlice is : []int which is different from the normal int type. in arrays its [5]int here its []int.

	// We add elements to the slice using the append function
	// Syntax : sliceName = append(sliceName, element(s))
	mySlice = append(mySlice, 6, 7, 8)                                     // This is paas by value. So the original slice is not modified.
	fmt.Println("After appending 6, 7, 8 value of mySlice is : ", mySlice) // The value of mySlice is : [1 2 3 4 5 6 7 8]
	fmt.Println("---------------------------")

	// To add elements to the slice at a specific index, we can use the append function with the spread operator.
	// Syntax : sliceName = append(sliceName[:index], appendValue, sliceName[index:]...)
	mySlice = append(mySlice[:5], append([]int{10, 11, 12}, mySlice[5:]...)...)
	fmt.Println("If we did `mySlice = append(mySlice[:5], append([]int{10, 11, 12}, mySlice[5:]...)...)` then the value of mySlice is : ", mySlice)
	// what this does is: mySlice[:5] gives [1 2 3 4 5] then we append another append of []int{10, 11, 12} with mySlice[5:] gives [10, 11, 12, 6, 7, 8]
	// which gives [1 2 3 4 5 10 11 12 6 7 8]
	fmt.Println("---------------------------")
	/* // This might create an issue Check later
	  var fruitsDB = []string{"Apple", "Orange", "Banana"}
		fruitsDB = append(fruitsDB, "Mango", "Pineapple", "Grapes")
		cutFruits := append(fruitsDB[1:3], fruitsDB[4:]...)
		fmt.Println("The value of fruitsDB is : ", fruitsDB)                                      // The value of fruitsDB is : [Apple Orange Banana Mango Pineapple Grapes]
		fmt.Println("We did `append(fruitsDB[1:3], fruitsDB[4:]...)` which gave us: ", cutFruits) // The cut fruits are : [Orange Banana Pineapple Grapes]
		fmt.Println("---------------------------")
	*/
	var fruitsDB = []string{"Apple", "Orange", "Banana"}
	fruitsDB = append(fruitsDB, "Mango", "Pineapple", "Grapes")
	fmt.Println("The value of fruitsDB is : ", fruitsDB) // The value of fruitsDB is : [Apple Orange Banana Mango Pineapple Grapes]
	cutFruits := append(fruitsDB[1:3], fruitsDB[4:]...)
	fmt.Println("We did `append(fruitsDB[1:3], fruitsDB[4:]...)` which gave us: ", cutFruits) // The cut fruits are : [Orange Banana Pineapple Grapes]
	fmt.Println("---------------------------")

	// Another way to create a slice is by using the make function
	// Syntax : sliceName := make([]dataType, length, capacity)

	var mySlice2 = make([]int, 5, 10)
	// It can accommodate five more elements without needing to allocate additional memory. When eleventh element is appended, the capacity is increased.
	fmt.Println("The value of mySlice2 is : ", mySlice2)       // The value of mySlice2 is : [0 0 0 0 0]
	fmt.Println("The length of mySlice2 is : ", len(mySlice2)) // The length of mySlice2 is : 5
	// or we could also do
	/*
		  var mySlice3 = make([]int, 5) //When an element is appended, the capacity is automatically increased (usually doubled).

			fmt.Println("The value of mySlice3 is : ", mySlice3)       // The value of mySlice3 is : [0 0 0 0 0]
			fmt.Println("The length of mySlice3 is : ", len(mySlice3)) // The length of mySlice3 is : 5
	*/
	fmt.Println("We update the value of mySlice2 by doing `mySlice2[2] = 200` etc.")
	mySlice2[0] = 100
	mySlice2[1] = 150
	mySlice2[2] = 200
	mySlice2[3] = 250
	mySlice2[4] = 300

	fmt.Println("This updates the value of the original slice mySlice to : ", mySlice2) // The value of mySlice is : [100 150 200 250 300]
	fmt.Println("---------------------------")

	// Sometimes we may need to copy the contents of one slice to another slice. We can use the copy function for this.
	// Syntax : copy(destinationSlice, sourceSlice)
	var mySlice3 = make([]int, 5)
	copy(mySlice3, mySlice2)
	fmt.Println("The value of mySlice3 after copying mySlice2 is : ", mySlice3) // The value of mySlice3 after copying mySlice2 is : [100 150 200 250 300]
	fmt.Println("---------------------------")

	// Sorting a slice
	// We can sort a slice using the sort package in Go. The sort package provides a function called Sort that can be used to sort a slice.
	// Syntax : sort.Sort(sliceName)
	// The sort.Sort function sorts the slice in increasing order.
	// To sort the slice in decreasing order, we can use the sort.Reverse function.
	// Syntax : sort.Sort(sort.Reverse(sliceName))
	// The sort.Reverse function returns a value that implements the sort.Interface interface, which is used by the sort.Sort function to sort the slice in decreasing order.

	mySlice4 := []int{5, 2, 6, 3, 1, 4}
	fmt.Println("The value of mySlice4 before sorting is : ", mySlice4) // The value of mySlice4 before sorting is : [5 2 6 3 1 4]
	sort.Ints(mySlice4)
	fmt.Println("The value of mySlice4 after sorting is : ", mySlice4) // The value of mySlice4 after sorting is : [1 2 3 4 5 6]
	sort.Sort(sort.Reverse(sort.IntSlice(mySlice4)))
	fmt.Println("The value of mySlice4 after sorting in reverse is : ", mySlice4) // The value of mySlice4 after sorting in reverse is : [6 5 4 3 2 1]
	fmt.Println("---------------------------")

	fmt.Println("How to remove a value from a slice")
	// To remove a value from a slice, we can use the append function with the spread operator.
	var coursesDB = []string{"Java", "JavaScript", "Go", "React", "PHP"}
	fmt.Println("We have the following in coursesDB : ", coursesDB) // We have the following in coursesDB :  [Java JavaScript Go React PHP]
	// To remove "React" from the slice which is on the index 3, we can use the following code
	// coursedDB := append(coursesDB[:3], coursesDB[4:]...)
	// fmt.Println("After removing React from coursesDB, we have : ", coursedDB) // After removing React from coursesDB, we have :  [Java JavaScript Go PHP]
	index := 3
	coursesDB = append(coursesDB[:index], coursesDB[index+1:]...)
	fmt.Println("After removing React from coursesDB, we have : ", coursesDB) // After removing React from coursesDB, we have :  [Java JavaScript Go PHP]

}
