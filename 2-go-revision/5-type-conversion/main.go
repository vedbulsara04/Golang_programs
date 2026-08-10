package main

import "fmt"

func main() {
	var i int = 42
	var f float64 = float64(i) / 10.0	// convert int to float

	var s string = fmt.Sprintf("%d", i)	// convert int to string

	fmt.Println(i, f)
	fmt.Println(s)
}
