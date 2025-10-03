package main

import "fmt"

func main(){
	
	// variable declared in main function
	message := "Hello from main func"
	fmt.Println(message)
	
	// fmt.Println(greeting) !Error: greeting is inaccessible!
	
	greet()	
}

func greet(){
	
	//variable declared in greet function
	greeting := "Greetings from greet func"
	fmt.Println(greeting)
	
	// fmt.Println(message) !Error: message is inaccessible!
}

