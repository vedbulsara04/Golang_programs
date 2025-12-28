package main

import "fmt"

func main() {
	var y = [3]int{1, 2, 3}
	
	y[0] = 10		// modify array
	fmt.Println(y[2])	// read array element
	fmt.Println(y)
}
