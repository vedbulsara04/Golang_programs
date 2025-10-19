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
		FirstName: "John",
		LastName: "Doe",
		Age: 30,
		email: "john@example.com",
	}
	
	// 3. Literal with Positional values (not recommended)
	p3 := Person {"Jane", "Smith", 28, "jane@example.com"}
	fmt.Printf("p3: %v", p3)
	
	// 4. Partial initialization
	p4 := Person{FirstName: "Alice", Age: 25}
	fmt.Printf("p4: %v", p4)

	// Accessing fields
	fmt.Printf(p2.FirstName)
	p2.Age = 31	// updates p2 age.
	
	// Pointers to structs
	p5 := &Person{FirstName: "Bob", Age: 40}
	fmt.Println(p5.FirstName) // Go's automatic dereferencing
	// Equivalent to: (*p5).FirstName
	
	// Embedded Structs
	e1 := Employee{
		Person: Person{
			FirstName: "Charlie",
			LastName: "Brown",
			Age: 35,
		},
		EmployeeID: 12345,
		Department: "Engineering",
	}
	
	// Accessing embedded fields
	fmt.Println(e1.FirstName) // Access through field promotion (shorter syntax)
	fmt.Println(e1.Person.FirstName) // Access through explicit path
	
	// Anonymous structs
	config := struct {
		Host string
		Port int
	}{
	
		Host: "localhost",
		Port: 8000,
	}
	fmt.Println(config)
}
