package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex // declaration of map.

func main() {
	m = make(map[string]Vertex) // initialization of map.
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])
}
