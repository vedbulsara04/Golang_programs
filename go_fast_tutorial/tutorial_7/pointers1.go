package main
import "fmt"

func main(){
	// Step 1: Create a pointer to int32 using new()
	var p *int32 = new(int32) // allocates memory, initializes to zero value (0)
	var i int32               // zero value for int32 is also 0
	
	// Step 2: Print the value p points to (will be 0, the zero value)
	fmt.Printf("The value p points to is: %v\n", *p) // prints: 0
	
	// Step 3: Print the value of i (this line has an error in original)
	fmt.Printf("The value of i is: %v\n", i) // should print i, not p
	
	// Step 4: Make p point to the memory address of variable i
	p = &i // now p points to i's memory location
	
	// Step 5: Change the value at the memory location p points to
	*p = 1 // this changes the value of i through the pointer
	
	// Step 6: Print both values - they should now be the same
	fmt.Printf("The value p points to is: %v\n", *p) // prints: 1
	fmt.Printf("The value of i is: %v\n", i)        // prints: 1
}
