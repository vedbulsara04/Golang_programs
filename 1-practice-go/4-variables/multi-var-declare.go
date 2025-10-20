package main

import "fmt"

func main(){
	// same type declaration
	var x, y, z int
	
	// different types with type inference
	var (
		name = "John Doe"
		age = 40
		salary = 100000.60
	)
	
	// using short declaration:
	a, b, c := 1, 2, 3
	firstName, lastName := "John", "Doe"
}

