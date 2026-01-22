package main

import "fmt"

func main(){
	// nil slice
	var x []int
	x = append(x, 10)
	fmt.Println(x)
	
	// append to slice with existing elements
	var y = []int{1, 2, 3}
	y = append(y, 4)
	fmt.Println(y)
	
	// append one slice to another using ... operator
	
	var a = []int{10, 20, 30}
	x = append(x, a...)
	fmt.Println(x)
}
