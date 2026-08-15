// Slice internals

package main

import "fmt"

func main() {
    s := make([]int, 3, 5)
    fmt.Printf("Len: %d, Cap: %d, Slice: %v\n", len(s), cap(s), s)
    // Len: 3, Cap: 5, Slice: [0 0 0]
    
    // Appending within capacity doesn't allocate
    s = append(s, 10)
    fmt.Printf("Len: %d, Cap: %d, Slice: %v\n", len(s), cap(s), s)
    // Len: 4, Cap: 5, Slice: [0 0 0 10]
    
    // Exceeding capacity triggers reallocation
    s = append(s, 20, 30)
    fmt.Printf("Len: %d, Cap: %d, Slice: %v\n", len(s), cap(s), s)
    // Len: 6, Cap: 10, Slice: [0 0 0 10 20 30] (capacity doubled)
}
