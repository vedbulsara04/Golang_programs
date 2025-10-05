// variable shadowing in Go

package main

import "fmt"

var x = "global" // Package level variable (global scope)

func main(){
	
	fmt.Println("value of global-scope x is printed: ", x) // accesses the global x
	
	x := "local" // new var in function scope (shadows global x)
	fmt.Println("value of function-level x is printed: ", x) // prints "local x"
	
	{
		x := "block" // new var in block scope (shadows function-level x)
		fmt.Println("value of block-level x is printed: ", x)
	}
	
	fmt.Println("value of function-level x is printed: ", x) // block level x is destroyed, function-level x is accessed
}

