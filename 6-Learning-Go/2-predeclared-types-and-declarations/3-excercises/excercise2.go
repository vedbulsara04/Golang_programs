/*
Write a program that declares a constant called value that can be 
assigned to both an integer and a floating-point variable. 
Assign it to an integer called i and a floating-point variable called f. 
Print out i and f.
*/

package main

import "fmt"

func main() {

	// declaring an untyped constant (flexibility)
	const value = 20
	
	i := value		// declaring int
	f := float64(value)	// declaring float
	
	fmt.Println("i: ", i)
	fmt.Println("f: ", f)
}
