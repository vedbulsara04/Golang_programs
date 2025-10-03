// Package Level Variable (Global Scope)

package main

import "fmt"

// must use 'var' cannot use ':='
var globalVar = "I am accessible everywhere in this package"

func main(){
	fmt.Println(globalVar)
}

func display(){
	fmt.Println(globalVar) // can access globalVar
}

