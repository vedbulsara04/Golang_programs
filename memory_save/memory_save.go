package main

import (
	"fmt"
	"runtime"
	"time"
)

func allocate() *[]byte {
	// Allocate 10MB and return pointer so GC can free it later
	data := make([]byte, 10*1024*1024)
	for i := range data {
		data[i] = byte(i % 256)
	}
	return &data
}

func main() {
	fmt.Println("Allocating 10MB...")
	buf := allocate()
	fmt.Printf("Allocated %d bytes\n", len(*buf))

	// Remove reference to allow GC to reclaim memory
	buf = nil

	// Force GC
	runtime.GC()
	fmt.Println("Memory should be freed. Check your process monitor!")

	// Wait so you can observe memory usage
	time.Sleep(5 * time.Second)
}