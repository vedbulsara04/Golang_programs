package main

import "fmt"

// Define an interface
type Shape interface {
	Area() float64
}

// Define a struct
type Circle struct {
	Radius float64
}

// Circle implements Shape interface
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

// Rectangle implements Shape interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Create values
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 3}

	// Use interface
	var s Shape

	s = c
	fmt.Println("Circle Area:", s.Area())

	s = r
	fmt.Println("Rectangle Area:", s.Area())
}
