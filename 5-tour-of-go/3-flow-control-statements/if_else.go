package main

import "fmt"

var age int = 18

func main() {
	if age < 20 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are an adult")
	}
}
