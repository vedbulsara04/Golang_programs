package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i // p is now 42
	fmt.Println(*p)

	*p = 21 // set value of i through pointer
	fmt.Println(i)

	p = &j       // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j)

}
