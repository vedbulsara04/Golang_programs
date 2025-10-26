package main

import "fmt"

func main() {
	// for range over a slice
	numbers := []int{10, 20, 30}
	fmt.Println("Iterating over slice: ")

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
