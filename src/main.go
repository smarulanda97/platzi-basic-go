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

	// Primitive types
	//
	// Integers:
	//
	// int = Depending of the OS (32 or 64 bits)
	// int8 = 8bits = -128 a 127
	// int16 = 16bits = -2^15 to 2^15-1
	// int32 = 32bits = -2^31 to 2^31-1
	// int64 = 64bits = -2^63 to 2^63-1
	//
	// Positive integers:
	//
	// uint8 = 8bits = 0 to 127
	// uint16 = 16bits = 0 to 2^16-1
	// uint32 = 32bits = 0 to 2^32-1
	// uint64 = 64bits = 0 to 2^64-1
	//
	// Decimals:
	//
	// float32 = 32 bits = +/- 1.18e^-38 to +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 to +/- -1.8e^308
	//
	// Text:
	//
	// string = "" Only double quotes
	//
	// Booleans:
	//
	// bool = true or false

	// fmt package
	//
	// Println: print with ln at the end of the line
	fmt.Println("This is a fmt.Println")
	// Printf:
	ageGate := 18
	message := "The minimum age to enter to this website is: "
	fmt.Printf("%s %d\n", message, ageGate)

	ageGateMessage := fmt.Sprintf("%s %d\n", message, ageGate)
	fmt.Println(ageGateMessage)
}
