/*
Every defer statement gets pushed onto a stack,
and when the function ends, they all pop off and execute from top to bottom (newest to oldest).
*/

package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

/*
defer fmt.Println(0) is added to the stack first (bottom)
defer fmt.Println(9) is added to the stack last (top)
When the function ends, Go starts executing from the top
So 9 prints first, then 8, then 7... all the way down to 0
*/
