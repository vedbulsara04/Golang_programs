package main

import "fmt"

func main(){
	// Array
	var intArr [3]int32 = [3]int32{1,2,3}
	// intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	fmt.Println(intArr)

	// Slice
	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)	
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	
	fmt.Println()
	
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	
	// Append multiple values with spread operator
	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
	
	// Another way to create Slice using make function
	var intSlice3 []int32 = make([]int32, 3, 8)
	// 3: specify length of slice, 8: specify capacity of slice.
	// By default capacity will be length of the slice
	fmt.Println(intSlice3)
	fmt.Printf("For intSlice3, the length is %v and capacity is %v\n", len(intSlice3), cap(intSlice3))
	
	
}

