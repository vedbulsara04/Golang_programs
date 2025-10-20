// Important string concepts

package main

import "fmt"

func main(){
	s := "Hello, 世界"
	
	// len() returns number of bytes, not characters
	fmt.Println("Number of bytes: ", len(s))	// 13
	
	// Counting runes (characters)
	fmt.Println("Counting runes: ", len([]rune(s)))	// 9

	// Indexing gives bytes
	fmt.Printf("Byte at index 7: %v\n", s[7])	// 228 (part of UTF-8 encoding)	
	
	// Use range to iterate over runes properly
	for i, r := range s{
		fmt.Printf("Position: %d, Rune: %c\n", i, r)
	}
}

