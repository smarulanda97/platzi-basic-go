package main

import (
	"fmt"
	"sync"
)

func printMessage(message string) {
	fmt.Println(message)
}

func calcSquareArea(sideSize int) (a, size int) {
	area := sideSize * sideSize
	return area, sideSize
}

type User struct {
	email    string
	username string
}

type Pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc Pc) getBrand() {
	fmt.Println(myPc.brand)
}

func (myPc *Pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

func (myPc Pc) String() string {
	return fmt.Sprintf("Brand: %s, Ram: %d, Disk: %d", myPc.brand, myPc.ram, myPc.disk)
}

type figure interface {
	area() float64
}

type square struct {
	base float64
}

func (s square) area() float64 {
	return s.base * s.base
}

type rectangle struct {
	base   float64
	height float64
}

func (r rectangle) area() float64 {
	return r.base * r.height
}

func calculate(f figure) {
	fmt.Println("Area: ", f.area())
}

func say(message string, wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println(message)
}

func say2(message string, c chan<- string) {
	c <- message
}

func message2(message string, c chan string) {
	c <- message
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
	// int8 = 8bits = 0 to 255
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

	// Structs
	user := User{email: "leidymejia17@hotmail.com", username: "leidymejia17"}
	fmt.Println(user)

	// Pointers
	num1 := 50
	num2 := &num1

	fmt.Println(num2)
	fmt.Println(*num2)

	myPc := Pc{ram: 16, disk: 320, brand: "MSI"}
	fmt.Println(myPc)
	myPc.getBrand()

	myPc.duplicateRam()
	fmt.Println(myPc)

	myPc.duplicateRam()
	fmt.Println(myPc)

	// Interfaces
	mySquare := square{base: 2}
	fmt.Println(mySquare)
	calculate(mySquare)

	myRectangle := rectangle{base: 2, height: 4}
	fmt.Println(myRectangle)
	calculate(myRectangle)

	// Interfaces list
	interfaceList := []interface{}{
		"Santiago", 24, 1,
	}
	fmt.Println(interfaceList...)

	// Go routines
	var wg sync.WaitGroup
	fmt.Println("Hello")

	wg.Add(2)
	go say("World", &wg)
	go func(message string) {
		wg.Done()
		fmt.Println(message)
	}("Bye bye.")
	wg.Wait()

	// Channels
	channel := make(chan string, 1)
	fmt.Println("My name is: ")

	go say2("Santiago Marulanda", channel)
	fmt.Println(<-channel)

	// Range, Close and Select in channels
	channel2 := make(chan string, 2)
	channel2 <- "Hello"
	channel2 <- "World"

	fmt.Println(len(channel2), cap(channel2))
	close(channel2)

	for message := range channel2 {
		fmt.Println(message)
	}

	email1 := make(chan string)
	email2 := make(chan string)

	go message2("Leidy", email1)
	go message2("Mejia", email2)

	for i := 0; i < 2; i++ {
		select {
		case message1 := <-email1:
			fmt.Println("Received message from email1: ", message1)
		case message2 := <-email2:
			fmt.Println("Received message from email2: ", message2)
		}
	}
}
