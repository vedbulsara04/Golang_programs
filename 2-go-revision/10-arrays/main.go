// Arrays: declaration, indexing, value semantics, comparison

package main

import "fmt"

func main() {
	var a [3]int = [3]int{1, 2, 3}
	b := a	// copy by value
	b[0] = 99

	fmt.Println("a:", a, "b:", b)

	// comparison [only arrays of comparable elements]
	fmt.Println("Equal? :", a == [3]int{1, 2, 3})
}
