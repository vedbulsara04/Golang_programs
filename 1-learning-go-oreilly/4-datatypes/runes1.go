package main

import "fmt"

func main(){
	var firstInitial rune = 'V' //good - type name matches usage.
	var lastInitial int32 = 'B' //bad - legal but confusing.
	
	fmt.Printf("First Initial: %v, Last Initial: %v", firstInitial, lastInitial)
}

