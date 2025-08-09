package main

import "fmt"

func main(){
	printMe()
	
	var printValue string = "Hello Ved!"
	printMe2(printValue)
	
	var numerator int = 11
	var denominator int = 2
	var result int = intDivision(numerator, denominator)
	fmt.Println("intDivision result: ", result)
}

func printMe(){
	fmt.Println("Go says hi!")
}

func printMe2(printValue string){
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) int{
	var result int = numerator/denominator
	return result
}


