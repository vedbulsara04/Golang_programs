package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30}

	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}
}
