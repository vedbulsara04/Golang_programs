package main

import "fmt"

func main(){
	
	// slice declaration
	var s1 []int	// nil slice
	fmt.Printf("s1: %v\n", s1)
	s2 := []int{}	// empty slice (not nill)
	fmt.Printf("s2: %v\n", s2)
	s3 := []int{1, 2, 3}	// slice literal
	fmt.Printf("s3: %v\n", s3)
	
	// using make
	s4 := make([]int, 5)	// length 5, capacity 5
	fmt.Printf("s4: %v\n", s4)
	s5 := make([]int, 3, 5)	// length 3, capacity 5 [0 0 0]
	fmt.Printf("s5: %v\n", s5)
	
	// Appending elements
	s3 = append(s3, 4, 5, 6)
	fmt.Printf("Appended s3: %v\n", s3)
	
	// Slicing operations
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr : %v\n", arr)
	slice1 := arr[1:4] // [2 3 4] (indices 1, 2, 3)
	slice2 := arr[:3] //
}

