// Maps are unordered collections of key-value pairs (hash tables)

package main

import "fmt"

func main(){
	
	// map declaration
	var m1 map[string]int	// nil map (can't add elements)
	fmt.Printf("m1: %v\n", m1)
	m2 := map[string]int{}	// empty map
 	fmt.Printf("m2: %v\n", m2)
	m3 := map[string]int{
		"Alice": 25,
		"Bob": 30,
	}
	fmt.Printf("m3: %v\n", m3)
	
	// using make
	m4 := make(map[string]int)
	m4["Charlie"] = 35
	fmt.Printf("m4: %v\n", m4)

}
		
