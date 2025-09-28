// var, const, int, float, string, type-casting, div & modulus, rune, bool.

package main

import "fmt"
import "unicode/utf8"

func main(){
	// var int:
	var intNum int16 = 35
	intNum = intNum + 1
	fmt.Println("intNum: ", intNum)
	
	// var float:
	var floatNum float32 = 12345678.9
	var floatNum2 float64 = 12345678.9
	fmt.Println("Float32: ", floatNum)
	fmt.Println("Float64: ", floatNum2)
	
	
	// Type casting:
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println("Type-casted Result: ", result)
	
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
	
	// no. of character is a string and not length:
	fmt.Println(len("test"))
	
	// Rune
	fmt.Println("RuneCountInString: ")
	fmt.Println(utf8.RuneCountInString("y"))
	
	var myRune rune = 'a'
	fmt.Println(myRune)
	
	// Bool
	var myBoolean bool = true
	fmt.Println(myBoolean)
	
	// Multi-variable initialization
	// var var1, var2 = 1, 2 [declared but unused]
	
	// Constants
	const myConstant string = "value1"
	// myConstant = "value2" [Error! It's a constant!]
	const pi float32 = 3.1415
}

