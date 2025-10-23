package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v1 := Vertex{1, 2}

	p := &v1  // p has value of v1
	p.X = 1e9 // 1000000000
	fmt.Println(v1)
}
