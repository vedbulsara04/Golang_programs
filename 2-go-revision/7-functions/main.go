// Functions : parameters, returns, multiple returns, named returns.
// Common pitfall: avoid naked returns in long functions (reduces readability).

package main

import "fmt"

// multiple returns

func divmod(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return	// named returns used
}

func add(a, b int) int { return a + b }

func main() {
	q, r := divmod(7, 3)
	fmt.Println("q, r: ", q, r)
	fmt.Println("add: ", add(2, 3))
}
