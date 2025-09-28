// Only simulates DB call dealy [no goroutines used]
// 1 id fetched at a random amount of time.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main(){
	t0 := time.Now()
	
	for i:=0; i<len(dbData); i++{
		dbCall(i) // go dbCall(i)
	}
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
}

func dbCall(i int){
	// Simulate the DB call dealy
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
}

