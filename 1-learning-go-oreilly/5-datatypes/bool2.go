package main

import "fmt"

func main(){
	var isLoggedIn bool = true
	var hasPermission bool = true
	var isBlocked bool = false
	
	// User can access if logged in AND has permission AND NOT blocked
	canAccess := isLoggedIn && hasPermission && !isBlocked
	fmt.Println(canAccess) //true
	
	// Show notification if logged in OR has important message
	// showNotification := isLoggedIn || hasImportantMessage
	
}
