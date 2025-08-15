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
	
	// error handling for non-existent names
	var age, ok = myMap2["Jason"]
	// delete(myMap2, "Adam") // delete a value from map
	if ok{
		fmt.Println("The age is %v", age)
	}else{
		fmt.Println("Invalid name")	
	}
	
	// Loops
	for name:= range myMap2{
		fmt.Printf("Name: %v\n", name)
	}
	
	for name, age := range myMap2{
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}
	
}

