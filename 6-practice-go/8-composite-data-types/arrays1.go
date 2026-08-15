package main

import "fmt"

func main(){
	
	// Array declaration
	var arr1 [5]int	// [0 0 0 0 0]
	fmt.Printf("arr1: %v\n", arr1)
	
	arr2 := [5]int{1, 2, 3, 4, 5}	// [1 2 3 4 5]
	fmt.Printf("arr2: %v\n", arr2)
	
	arr3 := [...]int{1, 2, 3}	// size inferred: [3]int
	fmt.Printf("arr3: %v\n", arr3)
	
	// Partial initialization
	arr4 := [5]int{1, 2}	// [1 2 0 0 0]
	fmt.Printf("arr4: %v\n", arr4)

	// Specific indices
	arr5 := [5]int{0:10, 4:50}
	fmt.Printf("arr5: %v\n", arr5)
	
	// Accessing elements
	fmt.Printf("access element at index 0: %v\n", arr2[0])
	
	arr2[0] = 100
	fmt.Printf("upated arr2 at index 0: %v", arr2)
	
	// Iterating over arrays
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("Index: %d, Value: %d\n", i, arr2[i])
	}
	
	// Using range
	fmt.Printf("\nUsing range: \n")
	for index, value := range arr2 {
		fmt.Printf("Index: %d, Value: %v\n", index, value)
	}
	
	// Arrays are value types
	arr6 := arr2	// Creates a copy
	arr6[0] = 999
	fmt.Println(arr2[0])
	fmt.Println(arr6[0])
	
}

