package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	p1 := Person{ // creating an instance of a Person.
		Name: "Ved",
		Age:  21,
		City: "Mumbai",
	}
	fmt.Printf("Person p1 = Name: %s, Age: %d, City: %s\n", p1.Name, p1.Age, p1.City)

	p2 := Person{"John Doe", 30, "Tokyo"} // create struct without naming fields.
	fmt.Printf("Person p2 = Name: %s, Age: %d, City: %s\n", p2.Name, p2.Age, p2.City)

	var p3 Person                              // zero-value struct.
	fmt.Printf("Zero value struct: %+v\n", p3) // %+v: It makes the output more readable by showing which value belongs to which field.
}
