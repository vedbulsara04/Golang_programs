package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i // The += operator means "add to the existing value."
	}
	fmt.Println(sum)
}
