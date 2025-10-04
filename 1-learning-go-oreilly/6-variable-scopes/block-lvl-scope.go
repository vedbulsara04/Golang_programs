package main

import "fmt"

func main(){
	x:= 10
	fmt.Println("x in main: ", x) // can access x here
	
	// if block
	if x > 5{
		y:= 20 // y only inside if block
		fmt.Println("x accessed by if block: ", x)
		fmt.Println("y accessed by if block: ", y)
	}
	
	// fmt.Println(y) // !Error: y is not accessible outside if block!
	
	// for loop block
	for i := 0; i < 3; i++ {
		z := 30 // z only exists inside for loop block
		fmt.Println("i:", i, "z", z)
	}
	
	// fmt.Println(i) // i is inaccessible (declared in for block)
	// fmt.Println(z) // z is inaccessible (declared in for block)
}
