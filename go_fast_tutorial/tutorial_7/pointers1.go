package main

import "fmt"

func main(){
	var p *int32 = new(int32) //initialize pointer
	var i int32
	fmt.Printf("The value p points to is: %v", *p) //reference the value of pointer
	fmt.Printf("The value of p is: %v", i)
	
	// *p = 10  //reference the value of pointer
	
}

