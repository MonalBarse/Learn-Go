package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go")
	// Arrays are a collection of elements that belong to the same type.
	// In Go, the size of an array is fixed. We can create an array by specifying the size of the array. Syntax:
	// var arrayName [size] dataType

	var myArray [5]int
	fmt.Println("The value of myArray is: ", myArray) // The value of myArray is: [0 0 0 0 0]
	// Since we did not assign any value to the array, the default value of all the elements in the array is 0.
	fmt.Println("---------------------------")
	// We can also assign values to the array while declaring it.
	var fruitList = [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("The value of fruitList is: ", fruitList)           // The value of fruitList is: [Apple Banana Orange Grapes]
	fmt.Println("The length of the fruitList is: ", len(fruitList)) // The length of the fruitList is: 4

	var vegList = [3]string{"Potato", "Tomato", "Onion"}
	fmt.Println("The value of vegList is: ", vegList)           // The value of vegList is: [Potato Tomato Onion]
	fmt.Println("The length of the vegList is: ", len(vegList)) // The length of the vegList is: 3
	fmt.Println("---------------------------")
	// We can also assign values to the array without specifying the size of the array.
	var cityList = [...]string{"Delhi", "Mumbai", "Bangalore", "Pune"}
	fmt.Println("The value of cityList is: ", cityList) // The value of cityList is: [Delhi Mumbai Bangalore Pune]
	fmt.Println("---------------------------")
	// We can also assign values to the array by specifying the index of the elements.
	var countryList = [5]string{1: "India", 3: "USA"}
	fmt.Println("The value of countryList is: ", countryList) // The value of countryList is: [ India  USA]
	fmt.Println("Notice that there is an empty sting at index 0 , 2 and 4.")
	fmt.Println("---------------------------")
	// Similarly if the array is of type int and we only assigned some of the values then the rest of the values will be 0.
	var numberList = [5]int{1: 10, 3: 30}
	fmt.Println("The value of numberList is: ", numberList) // The value of numberList is: [0 10 0 30 0]
	fmt.Println("---------------------------")

}
