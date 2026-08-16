package main

import(
	"fmt"
	"math"
)

func main() {
	// uint8 or byte - [example usage: raw pixel colour channel] (0 to 255)
	var red byte = 255
	var green byte = 128
	var blue byte = 0
	fmt.Printf("RGB: (%d, %d, %d)\n", red, green, blue)

	// uint16 - [example usage: network port number] (0 to 65535)
	var port uint16 = 8080
	fmt.Printf("Port number: %d\n", port)

	// uint32 - [example usage: IPv4 address as a packed integer] (0 to 4294967295)
	var ip uint32 = 0xC0A80101 // 192.168.1.1
	fmt.Printf("IP  : %d.%d.%d.%d\n", ip>>24, (ip>>16)&0xFF, (ip>>8)&0xFF, ip&0xFF)

	// uint64 - [example usage: file size in bytes, large counters] (0 to 8446744073709551615)
	var fileSize uint64 = 8_589_934_592 // 8 GB
	fmt.Printf("File: %d bytes = %.1f GB\n", fileSize, float64(fileSize)/1e9)

	// Boundaries
	fmt.Println("\n--- Boundaries ---")
	fmt.Println("uint8  : 0 ->", math.MaxUint8)
	fmt.Println("uint16 : 0 ->", math.MaxUint16)
	fmt.Println("uint32 : 0 ->", math.MaxUint32)
	fmt.Println("uint64 : 0 ->", uint64(math.MaxUint64))

	// Underflow Wrap-Around: Because u is an unsigned integer, 
	// it cannot represent negative numbers like -1. Subtracting 1 from 0 triggers an integer underflow, 
	// causing the value to wrap around to the maximum possible value for a uint8, 
	// which is 255.
	// Underflow wraps silently — classic unsigned gotcha
	fmt.Println("\n--- Underflow wrap (watch out!) ---")
	var u uint8 = 0
	u--
	fmt.Println("0 - 1 as uint8 =", u) // 255, not -1!

	// Safe subtraction - always check before subtracting
	a, b := uint8(10), uint8(20)
	if a >= b {
		fmt.Println("a - b: ", a - b)
	} else {
		fmt.Println("Cannot subtract: would underflow")
	}
}
