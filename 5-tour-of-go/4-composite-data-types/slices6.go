package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	fmt.Println("Original slice: ", numbers)

	numbers = append(numbers, 4)
	fmt.Println("Appending 4: ", numbers)

	numbers = append(numbers, 5, 6, 7)
	fmt.Println("Appending multiple numbers: ", numbers)

	moreNumbers := []int{8, 9, 10, 11}
	numbers = append(numbers, moreNumbers...) // ... unpacks the clice and is necessary.
	fmt.Println(numbers)
}
