package main

import (
	"fmt"
	"math"
)

func main() {
	// int8 - [example usage: Temperature value] (-128 to 127)
	var temp int8 = -45
	fmt.Printf("Temperature: %d\n", temp)

	// int16 - [example usage: Elevation in metres] (-32768 to 32767)
	var elevation int16 = -1120	// 1120 metres below sea level
	fmt.Printf("Elevation in metres: %d\n", elevation)

	// int32 or rune - Unicode code point (-2147483648 to 2147483647)
	var heart rune = '♥'	// rune literal is denoted in single quotes
	fmt.Printf("Rune: %v\n", heart)

	// int64 - large epoch timestamp
	var epoch int64 = 1_700_000_000	// '_' are the numeric separators
	fmt.Printf("Unix timestamp: %d\n", epoch)

	// int - platform-sized default (int64 vs int32 for 64 or 32 bit systems)
	score := 510
	fmt.Printf("Score: %d\n", score)

	// Boundaries
	fmt.Println("\n--- Boundaries ---")
	fmt.Println("int8: ", math.MinInt8, "-->", math.MaxInt8)
	fmt.Println("int16: ", math.MinInt16, "-->", math.MaxInt16)
	fmt.Println("int32: ", math.MinInt32, "-->", math.MaxInt32)
	fmt.Println("int64: ", math.MinInt64, "-->", math.MaxInt64)

	// Signed arithmetic - negatives work naturally
	fmt.Println("\n--- Signed arithmetic ---")
	var a, b int32 = 1_000_000, -250_000
	fmt.Println("a + b: ", a + b)
	fmt.Println("a - b: ", a - b)
}
