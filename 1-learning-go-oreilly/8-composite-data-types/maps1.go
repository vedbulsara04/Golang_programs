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
	
	// Adding/Updating elements
	m3["Charlie"] = 35	// adding
	fmt.Printf("m3: %v\n", m3)
	m3["Alice"] = 26	// updating
	fmt.Printf("m3: %v\n", m3)
	
	// accessing elements
	age := m3["Alice"]
	fmt.Printf("Alice's age: %v\n", age)	// 26
	
	// check if key exists
	age, exists := m3["Eve"]
	if exists {
		fmt.Println("Age: ", age)
	} else {
		fmt.Println("Key not found!")
	}
	
	// deleting elements
	delete(m3, "Bob")
	fmt.Printf("m3: %v\n", m3)
	
	// iterating over maps (the order is random!)
	for key, value := range m3 {
		fmt.Printf("%s: %d\n", key, value)
	}
	
	// Map length
	fmt.Println("Size of map m3: ", len(m3))
	
	// Maps are reference types
	// Both m3 and m5 point to the same map in memory.
	m5 := m3
	m5["Alice"] = 100
	fmt.Println(m3["Alice"]) // 100 (both point to the same map)
	
}

