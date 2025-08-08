// 

package main

import "fmt"

func main(){
	// var int
	var intNum int16 = 35
	intNum = intNum + 1
	fmt.Println("intNum: ", intNum)
	
	// float int
	var floatNum float32 = 12345678.9
	var floatNum2 float64 = 12345678.9
	fmt.Println("Float32: ", floatNum)
	fmt.Println("Float64: ", floatNum2)
	
	
	// Type casting
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)
	
	// Division & Modulus
	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println("Division: ", intNum1/intNum2)
	fmt.Println("Modulus: ", intNum1%intNum2) //To get remainder
	
	// String
	var myString string = "Hello_World"
	var myString2 string = "Hello \nWorld"
	var myString3 string = `Hello
				World`
	var myString4 string = "Hello" + "" + "World"
	fmt.Println(myString)
	fmt.Println(myString2)
	fmt.Println(myString3)
	fmt.Println(myString4)
	
	
	
	
}


