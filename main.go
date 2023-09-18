package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var age int32 = 37
	const isCool = true
	var size float32 = 2.3

	// Shorthand
	// name := "Brad"
	// email := "brad@gmail.com"

	name, email := "Sophie", "Sophie@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}