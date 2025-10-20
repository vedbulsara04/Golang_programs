package main

import "fmt"

func main(){
	var c64 complex64 = 1 + 2i
	var c128 complex128 = complex(3, 4)	// 3 + 4i
	
	// Extract real and imaginary parts
	real1 := real(c64)
	imag1 := imag(c64)
	
	real2 := real(c128)	// 3.0
	imag2 := imag(c128)	// 4.0
	
	fmt.Printf("Complex: %v, Real: %v, Imaginary: %v", c64, real1, imag1)
	fmt.Printf("\n")
	fmt.Printf("Complex: %v, Real: %v, Imaginary: %v", c128, real2, imag2)
	fmt.Printf("\n")
}

