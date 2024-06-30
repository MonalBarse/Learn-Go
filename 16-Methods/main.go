package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
	Salary    int
} // These fields are exported outside the package since they start with capital letter

func (u Employee) GetAge() { // This is a method with a receiver of type Employee
	fmt.Println("Is user an adult?")
	if u.Age >= 18 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func (u Employee) SalaryIncrement() {
	fmt.Println("Current salary of the employee is: ", u.Salary)

	newSalary := u.Salary + (u.Salary * 20 / 100)
	fmt.Println("New salary of the employee is: ", newSalary)
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

func main() {

	fmt.Println("Menhods in Golang")

	emp1 := Employee{"Son", "Goku", 130, 2341345}
	emp2 := Employee{"Monal", "Barse", 23, 998000}
	emp3 := Employee{"Vegeta", "Prince", 130, 8000000}
	emp4 := Employee{"Gohan", "Son", 130, 5000000}

	fmt.Printf(" Employee 1 %+v\n", emp1)
	fmt.Printf(" Employee 2 %+v\n", emp2)
	fmt.Printf(" Employee 3 %+v\n", emp3)
	fmt.Printf(" Employee 4 %+v\n", emp4)

	emp1.GetAge()
	emp2.SalaryIncrement()
	emp3.SalaryIncrement()
	emp4.SalaryIncrement()
	// These methods will not change the original struct values since they are not pointers
	// For them to do that, we need to use pointers
	// eg.
	// func (u *Employee) SalaryIncrement() {

	fmt.Printf(" Employee 1 %+v\n", emp1)
	fmt.Printf(" Employee 2 %+v\n", emp2)
	fmt.Printf(" Employee 3 %+v\n", emp3)
	fmt.Printf(" Employee 4 %+v\n", emp4)

}
