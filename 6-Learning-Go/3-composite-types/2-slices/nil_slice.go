package main

import "fmt"

func main() {
	var x = []int{1, 2, 3}
	var y []int	// nil slice
	
	fmt.Println(x)
	fmt.Println(y == nil)
}
