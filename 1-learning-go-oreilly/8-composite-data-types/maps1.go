// Maps are unordered collections of key-value pairs (hash tables)

package main

import "fmt"

func main(){
	
	// map declaration
	var m1 map[string]int	// nil map (can't add elements)
	m2 := map[string]int{}	// empty map
	m2 := map[string]int{
		"Alice": 25,
		"Bob": 30,
	}
}
