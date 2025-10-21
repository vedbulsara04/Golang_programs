// switch without a value

package main

import "fmt"

var score int = 85

func main() {
	// score := 85

	switch {
	case score >= 90:
		fmt.Println("Excellent!")
	case score >= 70:
		fmt.Println("Good Job!")
	case score >= 50:
		fmt.Println("You have Passed!")
	default:
		fmt.Println("You have Failed!")
	}
}
