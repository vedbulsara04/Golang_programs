// Strings in Go are immutable sequence of bytes (UTF-8 encoded)

package main

import (
	"fmt"
	"strings"
)

func main(){
	
	// string declaration
	var s1 string = "Hello World"
	fmt.Println(s1)
	s2 := "Hello, 世界"	// UTF-8 support
	fmt.Println(s2)
	
	// raw string literals (backticks)
	s3 := `This is a
	multiline
	string with \n not interpreted`
	fmt.Println(s3)
	
	// string length
	fmt.Println(len(s1))	// 13 (bytes, not characters!)
	
	// string indexing (gives bytes, not characters)
	fmt.Printf("First byte: %c\n", s1[0])	// 'H'
	
	// string concatenation
	fullName := "John" + " " + "Doe"
	fmt.Println(fullName)
	
	// string comparison
	fmt.Println("abc" == "abc")
	fmt.Println("abc" < "abd") // true (lexographic)
	
	// iterating over strings
	for i, char := range "Hello" {
		fmt.Printf("Index: %d, Character: %c, Unicode: %U\n", i, char, char)	
	}
	
	// common string operations
	
	fmt.Println(strings.Contains("seafood", "foo")) // true
	fmt.Println(strings.ToUpper("hello"))	// Hello
	fmt.Println(strings.Split("a,b,c", ",")) // [a b c]
	fmt.Println(strings.Split("hello", "")) // [h e l l o]			
}

