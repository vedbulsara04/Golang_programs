// Slices: append, len/cap, copy, nil vs empty, preallocation

package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	t := s[1:3]	// slicing
	s = append(s, 4)

	fmt.Println("s: ", s, "t: ", t)
	fmt.Println("len cap of s: ", len(s), cap(s))

	// preallocate
	p := make([]int, 0, 10)
	fmt.Println("p len cap:", len(p), cap(p))

    	// copy
    	dst := make([]int, len(s))
   	copy(dst, s)
	fmt.Println("dst:", dst)
}
