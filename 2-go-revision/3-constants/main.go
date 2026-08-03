package main

import "fmt"

const Pi = 3.14159	// untyped constant
const (
	A int = 10	// typed constant
	B     = 20	// untyped constant, becomes appropriate when used
)

func main() {
	var r int = 2

	// untyped constant Pi can be used in float64 context
	area := Pi * float64(r*r)
	fmt.Println("A, B:", A, B, "Area:", area)
}
