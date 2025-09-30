package main

import "fmt"

func main(){
	var x int = 10
	var y byte = 100 //byte = uint8
	
	var sum1 byte = byte(x) + y
	var sum2 int = x + int(y)
	
	fmt.Printf("int to byte: %v, byte to int: %v", sum1, sum2)
}

