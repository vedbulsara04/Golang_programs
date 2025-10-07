// Strings in Go are immutable sequence of bytes (UTF-8 encoded)

package main

import (
	"fmt"
	"strings"
)

func main(){
	
	// string declaration
	var s1 string = "Hello World"
	s2 := "Hello, 世界"	// UTF-8 support
	
	// raw string literals (backticks)
	s3 := `This is a
	multiline
	string with \n not interpreted`
	
	// string length
	fmt.Println(len(s1))	// 13 (bytes, not characters!)
	
	// string indexing (gives bytes, not characters)
	fmt.Printf("First byte: %c\n", s1[0])	// 'H'
	
	// string concatenation
	fullName := "John" + " " + "Doe"
	
	// string comparison
	fmt.Println("abc" == "abc")
	fmt.Println("abc" < "abd") // true (lexographic)
	
	// iterating over strings
	for i, char := range 'Hello' {
		fmt.Printf("Index: %d, Character: %c, Unicode: %U\n", i, char, char)	
	}
	
	
}
