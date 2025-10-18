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
	
}
