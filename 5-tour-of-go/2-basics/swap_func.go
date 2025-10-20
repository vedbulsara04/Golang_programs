// In Go, functions can return any number of results.

package main

import "fmt"

// Note; swap is a custom function and not a built-in function of Go.
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
