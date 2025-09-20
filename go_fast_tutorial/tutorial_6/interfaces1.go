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

func canMakeIt(e engine, kms uint16){
	if kms<=e.kmsleft(){
		fmt.Println("You can make it!")
	}else{
		fmt.Println("Need to fuel up!")
	}
}

func main(){
	var myICE icEngine = icEngine{25, 15}
	canMakeIt(myICE, 120) 
}

