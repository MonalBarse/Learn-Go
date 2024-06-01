package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {

	fmt.Println("Structs in Golang")
	// Structs are similar to classes in other programming languages
	// Structs are used to create user defined data types
	// Syntax : type struct_name struct { key1 data_type1 key2 data_type2 }

	// There is no inheritance in Golang, but we can achieve composition by embedding one struct into another struct
	// No super or parent keyword in Golang
	// Creating a struct
	emp1 := Employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
		salary:    600000,
	} // or emp1 := Employee{"Sam", "Anderson", 55, 6000} since we know the order of the fields
	emp2 := Employee{"Monal", "Barse", 29, 998000}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	fmt.Println("A more nicer syntax for printing the struct is: % +v ")
	fmt.Printf(" Employee 1's details are: %+v\n", emp1)
	fmt.Printf(" Employee 2's details are: %+v\n", emp2)

}

/*
  This is only basics of structs remember to explore other uses too.
*/
