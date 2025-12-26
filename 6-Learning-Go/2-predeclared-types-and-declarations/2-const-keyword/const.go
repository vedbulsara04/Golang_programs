package main

import "fmt"

const x int64 = 10

func main() {
	const y = "hello"
	
	fmt.Println(x)
	fmt.Println(y)
	
	x = x + 1	// will not compile!
	y = "bye"	// will not compile!
	
	fmt.Println(x)
	fmt.Println(y)
}
