/*
defer is a keyword in Go that delays the execution of a function until the surrounding function finishes.
Think of it like saying: "Do this... but wait until everything else is done first."
*/

package main

import "fmt"

func main() {
	defer fmt.Println("World")

	fmt.Println("Hello")
}
