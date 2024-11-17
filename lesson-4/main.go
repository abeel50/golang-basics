package main

import "fmt"

func main() {
	age := 20
	name := "John"
	//print
	fmt.Print("Hello, World!")
	fmt.Print("Hello, World!\n")
	fmt.Print("Hello\n")

	fmt.Println("My name is", name, "and I am", age, "years old.")

	// formatted print
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// save formatted string
	var st = fmt.Sprintf("My name is %s and I am %d years old.\n", name, age)
	fmt.Println(st)
}
