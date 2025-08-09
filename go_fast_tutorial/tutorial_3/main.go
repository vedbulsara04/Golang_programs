package main

import "fmt"

func main(){
	printMe()
	
	var printValue string = "Hello Ved!"
	printMe2(printValue)
	
	var numerator int = 11
	var denominator int = 2
	var result, remainder int = intDivision(numerator, denominator)
	// fmt.Println("intDivision result: ", result)
	fmt.Printf("The result of integer division is %v with remainder as %v", result, remainder)
}

func printMe(){
	fmt.Println("Go says hi!")
}

func printMe2(printValue string){
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int) {
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder
}


