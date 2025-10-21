package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 { // like a while loop (loops until condition is satisfied)
		sum += sum // equivalent to sum = sum + sum
	}
	fmt.Println(sum)
}
