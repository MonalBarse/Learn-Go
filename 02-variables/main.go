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
}

// Placeholders in Printf in alphabetical order

// %b : binary values          : eg:- "%b" for 24 will give 11000
// %c : character values       : for strings it will take the first character eg:- "%c" will give "H" for "Hello"
// %d : decimal values         : eg:- "%d" will give 5 for 5
// %f : floating point values  : eg:- "%f" will give 5.500000 for 5.5
// %o : octal values           : eg:- "%o" for 234 will give 352
// %p : pointer values         : eg:- "%p" for &x will give 0xc0000b6010
// %s : string values          : eg:- "%s" for "Hello" will give Hello
// %t : boolean values         : eg:- "%t" for false will give false
// %v : default format         : eg:- "%v" for 5 will give 5
// %x : hexadecimal values     : eg:- "%x" for 324324 will give 4f9f4

// %+v : print the value in a default format with field names : eg in a struct say {name: "John", age: 25} for %+v will give {name: John, age: 25}

// capital letter placeholder (also in alphabetical order)

// %E : scientific notation
// %F : decimal point but no exponent
// %G : %E for large exponents, %F otherwise
// %S : the value in a default format
// %T : the type of the value
// %U : Unicode format
// %X : base 16, upper-case letters

// Walrus Operator
/*
  The walrus operator := is used to declare and initialize a variable in a single line.
  It is shorthand for declaring a variable and then assigning a value to it.

  we normally do this: var x int = 5; which is a complete statement in itself
  but we can also do this: x := 5; which is a shorthand for the above statement

  We can also use the walrus operator to declare multiple variables in a single line.
  For example, x, y := 5, "Hello" will declare x as an integer with value 5 and y as a string with value "Hello".
  It automatically assigns the type of the variable based on the value assigned to it.

  The walrus operator can only be used inside functions and not at the package level.
*/
