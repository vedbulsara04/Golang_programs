package main

import "fmt"

type icEngine struct{
	kml uint16
	litres uint16
}

func (e icEngine) kmsleft() uint16{
	return e.litres*e.kml
}

func main(){
	var myICE icEngine = icEngine{25, 15}
	// 25*15 gives integer overflow error due to uint8 taking only upto 0-255 value. [performs wrong calculation]
	fmt.Printf("Total kms left in tank: %v", myICE.kmsleft()) 
}

