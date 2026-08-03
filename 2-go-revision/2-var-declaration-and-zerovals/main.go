package main

import "fmt"

func main() {
	// short variable declaration
	x := 42
	// var with explicit type
	var y int
	// multiple variable declaration
	var a, b = "foo", 3.14
	// zero value for bool is false, for int is 0, for string is ""
	var z bool

	fmt.Printf("x=%d y=%d a=%s b=%f z=%v\n", x, y, a, b, z)
}

