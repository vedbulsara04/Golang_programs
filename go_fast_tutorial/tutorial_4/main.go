package main

import "fmt"

func main(){
	var intArr [3]int32 = [3]int32{1,2,3}
	// intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	fmt.Println(intArr)
}

