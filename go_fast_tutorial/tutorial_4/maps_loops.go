package main

import "fmt"

func main() {
	// Maps in Go

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap) // o.p : map[]

	// directly declare a map
	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2)
	fmt.Println("Age of Adam:", myMap2["Adam"])

	// error handling for non-existent names
	var age, ok = myMap2["Jason"]
	// delete(myMap2, "Adam") // delete a value from map
	if ok {
		fmt.Println("The age is %v", age)
	} else {
		fmt.Println("Invalid name")
	}

	// Loops
	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	// loop over arrays and slices, i=index | v=value
	var intArr [3]int32 = [3]int32{2, 3, 4}
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	// While loop in go using for loops

	/*
		var i int = 0
		for i<10{
			fmt.Println(i)
			i = i+1
		}
		fmt.Println()
	*/

	/*
		var i int = 0
		for{
			if i >= 10{
				break
			}
			fmt.Println(i)
			i = i + 1
		}
	*/

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
