package main

import "fmt"

type icEngine struct{
	kml uint16
	litres uint16
}

type electricPU struct{
	kmpkwh uint16
	kwh uint16 
}

func (e icEngine) kmsleft() uint16{
	return e.litres*e.kml
}

func (e electricPU) kmsleft() uint16{
	return e.kwh*e.kmpkwh
}

type engine interface{
	kmsleft() uint16
}


func main(){
	var myICE icEngine = icEngine{25, 15}

	fmt.Printf("Total kms left in tank: %v", myICE.kmsleft()) 
}

