package main

import "fmt"

func main(){
	// Maps in Go
	
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap) // o.p : map[]
	
	// directly declare a map
	var myMap2 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Println(myMap2)
	fmt.Println("Age of Adam:", myMap2["Adam"])
}

