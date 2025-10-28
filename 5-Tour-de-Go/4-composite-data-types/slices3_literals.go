/*
A slice literal is like an array literal without the length.
It is a concise way to create and initialize a slice
*/

package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
	}
	fmt.Println(s)
}
