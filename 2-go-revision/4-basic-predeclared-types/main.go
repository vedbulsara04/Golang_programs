package main

import "fmt"

func main() {
	var i int = 7
	var f float64 = 4.3
	var s string = "Gopher"
	var b byte = 'A'	// alias for uint8
	var r rune = '✓'	// alias for int32, holds unicode code point
	var c complex128 = complex(1, 2)

	fmt.Println(i, f, s, b, r, c)
        fmt.Printf("rune=%c\n", r)
}
