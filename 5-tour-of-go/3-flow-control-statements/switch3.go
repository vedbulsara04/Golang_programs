package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday() // time.Now() gets the current date and time | .Weekday() extracts just the day of the week (returns a number: Sunday=0, Monday=1, Tuesday=2, etc.)

	switch time.Saturday { // Starts a switch statement that checks against time.Saturday which has the value 6
	case today + 0:
		fmt.Println("Today is Saturday")
	case today + 1:
		fmt.Println("Tomorrow is Saturday")
	case today + 2:
		fmt.Println("In two days..")
	case today + 3:
		fmt.Println("In 3 days, today is Wednesday..")
	case today + 4:
		fmt.Println("In 4 days, today is Tuesday..")
	default:
		fmt.Println("Too far away...")
	}
}
