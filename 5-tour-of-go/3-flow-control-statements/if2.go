// Go allows you to execute a short statement before condition

package main

import "fmt"

func main() {
	if num := 10; num > 5 {
		fmt.Println("num is greater than 5")
	}
	// Note: num is part of if block
}
