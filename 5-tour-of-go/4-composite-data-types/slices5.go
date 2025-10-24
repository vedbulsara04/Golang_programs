package main

import "fmt"

func main() {
	a := make([]int, 5) // slice with length & capacity = 5
	fmt.Println(a)

	b := make([]int, 0, 5) // slice with length = 0 & capacity = 5
	fmt.Println(b)

	c := b[:2] // creates a slice with length = 2 and capacity = 5
	fmt.Println(c)

	d := c[2:5] // creates a slice with length = 3 and capacity = 3
	fmt.Println(d)
}
