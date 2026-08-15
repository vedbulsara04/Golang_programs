package main

import "fmt"

func main(){
	var f32 float32 = 3.14159265359
	var f64 float64 = 3.14159265359		// more precision
	
	// Scientfic notation
	var bigNum = 1.23e9	// 1.23 * 10^9
	var smallNum = 1.23e-9	// 1.23 * 10^-9
	
	fmt.Printf("float32: %.10f\n", f32)	// limited precision
	fmt.Printf("float64: %.10f\n", f64)	// better precision
	fmt.Printf("bigNum: %.2f\n", bigNum)
	fmt.Printf("smallNum: %.10f\n", smallNum)
	fmt.Printf("bigNum in scientific notation: %e\n", bigNum)
	fmt.Printf("smallNum in scientific notation: %e\n", smallNum)
}

