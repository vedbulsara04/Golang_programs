package main

import "fmt"
import "strings"

func main(){

	// string
	var myString = "résumé"
	var indexed = myString[0]
	
	fmt.Printf("%v, %T\n", indexed, indexed)
	
	for i, v := range myString{
		fmt.Println(i, v)
	}
	
	fmt.Printf("\nThe lenght of 'myString' is: %v\n", len(myString))
	
	// runes
	var myRune = 'a'
	fmt.Printf("\nmyRune = %v\n", myRune)

	// string building using concatenation
	var strSlice= []string{"H","a","m","i","l","t","o","n","-","4","4"}
	var concatStr = ""
	for i := range strSlice{
		concatStr += strSlice[i]
	}
	fmt.Printf("\n%v\n", concatStr)
	
	//string building using Go's strings package
	
	
}

