package main

import "fmt"

func main() {
	// Declare a variable.
	number := 42

	// Create a pointer to `number` using the '&' operator.
	// `ptr` now holds the memory address of `number`.
	ptr := &number

	fmt.Println("Initial value of number:", number)
	fmt.Println("Memory address of number:", ptr)

	// Dereference the pointer using the '*' operator to change the value.
	// This modifies the original `number` variable.
	*ptr = 100

	fmt.Println("New value of number:", number)
	fmt.Println("New value of dereferenced pointer:", *ptr)
}
