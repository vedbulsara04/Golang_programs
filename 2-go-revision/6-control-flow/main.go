// Control FLow: if, else, switch, defer.

package main

import "fmt"

func main() {

	x := 3
	if x%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

	fmt.Println("For Loop:")
	for i := 0; i < 3; i ++ {
		fmt.Println(i)
	}

	fmt.Println("Switch: ")
	switch x {
	case 1:
		fmt.Println("One")
	case 2, 3:
		fmt.Println("Two or Three")
	default:
		fmt.Println("Other")
	}

	defer fmt.Println("deferred: runs at the end of function")
	fmt.Println("End of Main")
}
