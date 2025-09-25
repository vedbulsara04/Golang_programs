// Simulates DB call dealy [goroutines have been used]
// fetches all ids using goroutines.

package main

import (
	"fmt" // For printing to console (Printf, Println)
	"math/rand" // For generating random numbers (Float32())
	"time" // For time operations (Now, Sleep, Duration, Since)
	"sync" // For synchronization primitives (WaitGroup)
)

var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main(){
	t0 := time.Now()
	
	for i:=0; i<len(dbData); i++{
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
}

func dbCall(i int){
	// Simulate the DB call dealy
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	wg.Done()
}

