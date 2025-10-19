// Structs are composite types that group fields of different types.

package main

import "fmt"

// Struct definition
type Person struct {
	FirstName string
	LastName string
	Age int
	email string
}

type Employee struct {
	Person string // (Go's way of composition: [embedded Person struct])
	EmployeeID int
	Department string
}

func main(){
	// Struct initialization methods
	
	// 1. Zero Value | Declares a Person variable p1 using zero value initialization
	// All fields will be set to their zero values (empty strings for strings, 0 for int)
	var p1 Person
	fmt.Println(p1)
	
	
	// 2. Literal with field names (recommended method) [struct literal syntax]
	p2 := Person {
		FirsName: "John",
		LastName: "Doe",
		Age: 30,
		Email: "john@example.com",
	}
	
	// 3. Literal with Positional values (not recommended)
	p3 := Person {"Jane", "Smith", 28, "jane@example.com"}
	
	// 4. Partial initialization
	p4 := Person{FirstName: "Alice", Age: 25}

	// Accessing fields
	fmt.Printf(p2.FirstName)
	p2.Age = 31	// updates p2 age.	
}
