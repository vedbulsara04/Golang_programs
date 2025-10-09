package main

import "fmt"

func main(){
	
	// Array declaration
	var arr1 [5]int	// [0 0 0 0 0]
	fmt.Printf("arr1: %v", arr1)
	
	arr2 := [5]int{1, 2, 3, 4, 5}	// [1 2 3 4 5]
	fmt.Printf("arr2: %v", arr2)
	
	arr3 := [...]int{1, 2, 3}	// size inferred: [3]int
	
	// Partial initialization
	arr4 := [5]int{1, 2}	// [1 2 0 0 0]
}

