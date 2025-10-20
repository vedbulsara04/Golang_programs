/*

The blank identifier (_) is a special identifier in Go that lets you ignore/discard values you don't need. It's like a "trash bin" for unwanted values.

Why use it: In Go, if you declare a variable but don't use it, the compiler throws an error. The blank identifier solves this problem.

*/

package main

import "fmt"

func getNumbers() (int, int) {
	return 42, 100
}

func main() {
	// Only want the first number, ignore the second
	first, _ := getNumbers()
	fmt.Println("First number:", first)
	
	// Ignore index in loop
	colors := []string{"Red", "Green", "Blue"}
	for _, color := range colors {
		fmt.Println(color)
	}
}

