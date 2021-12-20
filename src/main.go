package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	// Constants
	const PI float64 = 3.14
	const PI2 = 3.14

	fmt.Println("PI: ", PI)
	fmt.Println("PI2: ", PI2)

	// Integers
	var age int = 24
	fmt.Println("Your age is: ", age)

	// Default values for types
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("Default for int: ", a, "Default for float: ", b, "Default for string: ", c, "Default for bool: ", d)

	// Other way to define variables
	x := "Other way to define variables is using ':='"
	fmt.Println(x)
}
