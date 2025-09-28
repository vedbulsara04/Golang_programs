package main

import (
	"fmt"
	"errors"
)

func main(){
	printMe()
	
	var printValue string = "Hello Ved!"
	printMe2(printValue)
	
	var numerator int = 10
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)
	// check if error encountered:
	/* if err!=nil{
		fmt.Printf(err.Error())
	}else if remainder == 0{
		fmt.Printf("The result of integer division is %v", result)
	}else{
		// fmt.Println("intDivision result: ", result)
		fmt.Printf("The result of integer division is %v with remainder as %v", result, remainder)
	}*/
	
	switch{
		case err != nil:
			fmt.Printf(err.Error())
		case remainder == 0:
			fmt.Printf("The result of integer division is %v", result)
		default:
			fmt.Printf("The result of integer division is %v with remainder as %v", result, remainder)
	}
	
	switch remainder{
		case 0:
			fmt.Printf("The division was exact")
		case 1,2: 
			fmt.Printf("The division was close")
		default:
			fmt.Printf("The division was not close")
	}
}

func printMe(){
	fmt.Println("Go says hi!")
}

func printMe2(printValue string){
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator==0{
		err = errors.New("Cannot divide by zero!")
		return 0, 0, err
	}
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
}

