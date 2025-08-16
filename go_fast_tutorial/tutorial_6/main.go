package main

import "fmt"

// creating a struct
type icEngine struct{
	kml uint8 //mileage | o.p(undeclared values): 0 [uint8]
	litres uint8 //litres of fuel left | o.p(undeclared values): 0 [uint8]
}

func main(){
	//var myICE icEngine
	//var myICE icEngine = icEngine{kml:25, litres:15}
	//var myICE icEngine = icEngine{25, 15}
	var myICE icEngine
	myICE.kml = 20
	myICE.litres = 9
	
	fmt.Println(myICE.kml, myICE.litres)	
	
}

