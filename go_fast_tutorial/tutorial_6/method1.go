package main

import "fmt"

type icEngine struct{
	kml uint8
	litres uint8
}

func (e icEngine) kmsleft() uint8{
	return e.litres*e.kml
}

func main(){
	var myICE icEngine = icEngine{25, 15}
	fmt.Printf("Total kms left in tank: %v", myICE.kmsleft()) 
}

