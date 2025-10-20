package main

import (
	"fmt"
	"math"
)

func main(){
	
	// Signed Integers
	var i8 int8 = 127		// -128 to 127
	var i16 int16 = 32767           // -32768 to 32767
	var i32 int32 = 2147483647      // -2^31 to 2^31-1
	var i64 int64 = 9223372036854775807  // -2^63 to 2^63-1

	// Unsigned Integers
	var ui8 uint8 = 255		// 0 to 255
	var ui16 uint16 = 65535		// 0 to 65535
	var ui32 uint32 = 4294967295	// 0 to 2^32-1
	
	// Platform-dependent int (32 or 64 bit)
	var i int = 42
	
	// Special aliases
	var b byte = 255 		// (alias for uint8)
	var r rune = '世'		// (alias for int32, represents Unicode code point)
	
	fmt.Println("Signed Integers:")
	fmt.Printf("int8: %d, min: %d, max: %d\n", i8, math.MinInt8 ,math.MaxInt8)
	fmt.Printf("int16: %d, max: %d\n", i16, math.MaxInt16)
	fmt.Printf("int32: %d, max: %d\n", i32, math.MaxInt32)
	fmt.Printf("int64: %d, max: %d\n", i64, math.MaxInt64)
	fmt.Println("")
	
	fmt.Println("Unsigned Integers:")
	fmt.Printf("uint8: %d, max: %d\n", ui8, math.MaxUint8)
	fmt.Printf("uint16: %d, max: %d\n", ui16, math.MaxUint16)
	fmt.Printf("uint32: %d, max: %d\n", ui32, math.MaxUint32)
	fmt.Println("")
	
	fmt.Printf("Platform-dependent int: %v\n", i)
	fmt.Println("")
	
	fmt.Printf("byte(alias for uint8): %v\n", b)
	fmt.Println("")
	
	fmt.Printf("rune: %d, character: %c\n", r, r)
}

