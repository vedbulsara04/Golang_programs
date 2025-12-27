/*
WAP that declares an integer variable called i with the value 20.
Assign i to a floating point variable named f. Print out i and f.
*/

package main

import "fmt"

func main() {
	var i int = 20
	// i := 20
	
	f := float64(i)
	
	fmt.Println(i, f)
}
