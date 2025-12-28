package main

import "fmt"

func main() {
	var x = [3]int{1, 2, 3}
	var y = [4]int{1, 2, 3}	// Prints as {1, 2, 3, 0}
	
	fmt.Println("Print arrays: ")
	fmt.Println(x)	
	fmt.Println(y)
	fmt.Println("use len func:")
	fmt.Println(len(x))
	fmt.Println(len(y))		
}
