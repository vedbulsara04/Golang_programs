package main

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message, i)
		time.Sleep(500 * time.Millisecond) // Simulate some work
	}
}

func main() {
	// Start a goroutine
	go printMessage("From goroutine")

	// Run this in the main thread
	printMessage("From main function")

	// Give goroutine some time to finish before program exits
	time.Sleep(3 * time.Second)
}
