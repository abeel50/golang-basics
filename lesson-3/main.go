package main

import "fmt"

func main() {
	fmt.Println("-------------Strings--------------")
	// strings
	var firstName string = "John"
	var lastName = "Doe"
	var anotherName string //Default empty string
	fmt.Println(firstName, lastName, anotherName)

	firstName = "Jane"
	anotherName = "Smith"

	fmt.Println(firstName, lastName, anotherName)

	//another way of initializing a variable without var
	//! this can't be used outside a function
	name := "John Doe"
	fmt.Println(name)

	// Integers
	fmt.Println("---------Integers---------")
	var age int = 40
	var anotherAge = 30 // type inference
	yetAnotherAge := 45

	fmt.Println(age, anotherAge, yetAnotherAge)

	// Bits & Memory
	fmt.Println("---------Bits & Memory---------")
	var num1 int8 = 25   // 8 bits -127 to 127
	var num2 int16 = 256 //16 bits -32768 to 32767
	var num3 int32 = 256 //32 bits -2147483648 to 2147483647
	var num4 int64 = 256 //64 bits -9223372036854775808 to 9223372036854775807

	fmt.Println(num1, num2, num3, num4)

	fmt.Println("---------Floats---------")
	var num5 float32 = 3.14
	var num6 float64 = 3.14

	fmt.Println(num5, num6)
}
