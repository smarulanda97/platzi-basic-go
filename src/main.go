package main

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

func calcSquareArea(sideSize int) (a, size int) {
	area := sideSize * sideSize
	return area, sideSize
}

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

	// Functions
	printMessage("This print is by function printMessage()")

	var area, size int = calcSquareArea(5)
	fmt.Printf("The area of the square with side size (%d) is: %d\n", size, area)

	var area2, _ int = calcSquareArea(6)
	fmt.Printf("The area of the square with side size (any) is: %d\n", area2)

	// Loops
	//
	// Basic for:
	//
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//
	// For while:
	//
	counter := 0
	for counter < 5 {
		fmt.Println(counter)
		counter++
	}

	// Arrays and Slices
	//
	// Arrays:
	//
	var users [4]string
	users[0] = "smarulanda9703@gmail.com"
	users[1] = "santiagomm1997@gmail.com"

	fmt.Println(users, len(users), cap(users))
	//
	// Slices:
	//
	users2 := []string{"smarulanda97@outlook.com", "leidymejia17@hotmail.com", "d.ana26@hotmail.com", "frasul301@hotmail.com"}
	fmt.Println(users2, len(users2), cap(users2))
	fmt.Println(users2[0])
	fmt.Println(users2[:2])
	fmt.Println(users2[2:3])
	fmt.Println(users2[3:])

	users2 = append(users2, "test@gmail.com")
	fmt.Println(users2)

	users3 := []string{"smarulanda9703@gmail.com", "santiagomm1997@gmail.com"}
	users3 = append(users3, users2...)
	fmt.Println(users3)

	for i, valor := range users3 {
		fmt.Printf("users3[%d]=%s\n", i, valor)
	}
	// If index is not necessary
	for _, valor := range users3 {
		fmt.Printf("users3=%s\n", valor)
	}

	// Maps
	userData := make(map[string]string)
	userData["first_name"] = "Santiago"
	userData["last_name"] = "Marulanda"

	fmt.Println(userData)
	for i, value := range userData {
		fmt.Println(i, value)
	}

	fmt.Println(userData["first_name"])
	firstName, ok := userData["first_name"]

	fmt.Println(firstName, ok)
}
