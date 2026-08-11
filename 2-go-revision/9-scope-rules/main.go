// Scope Rules: package/function/block-scope, shadowing, blank identifier _
// Pitfall: shadowing can cause subtle bugs when reusing variable names

package main

import "fmt"

var g = "package-level"

func main() {
	g := "shadowed"	// shadows package-level g in this block
	fmt.Println("g:", g)

	_, y := foo()	// ignore first return using blank identifier
	fmt.Println("y:", y)
}

func foo() (int, int) {
	return 1, 2
}


