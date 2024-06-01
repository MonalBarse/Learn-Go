package main

import "fmt"

func main() {
	var x int = 5
	fmt.Printf("x is an integer with value: %v\n", x)

	var y float64 = 5.5
	fmt.Printf("y is a float with value: %v\n", y)

	var z string = "Hello, World!"
	fmt.Printf("z is a string with value: %v\n", z)

	var a bool = true
	fmt.Printf("a is a boolean with value: %v\n", a)

	var j complex128 = 5 + 5i
	fmt.Printf("j is a complex number with value: %v\n", j)

	var b byte = 'a'
	fmt.Printf("b is a byte with value: %v\n", b)

	var c rune = 'a'
	fmt.Printf("c is a rune with value: %v\n", c)
	// a rune is an alias for int32 and is used to distinguish character values from integer values.

	var d uint = 5
	fmt.Printf("d is an unsigned integer with value: %v\n", d)

	var sum complex128 = j + 4*j
	fmt.Printf("sum is a complex number with value: %v\n", sum)

	// we can type convert variables from higher to lower precision types for eg. we can convert a complex to int but not vice versa
	var sum2 int = int(sum)
}
