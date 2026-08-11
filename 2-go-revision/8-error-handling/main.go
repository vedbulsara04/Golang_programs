// Error Handling: error, errors.New, branching on errors.

package main

import (
	"errors"
	"fmt"
)

func safeSqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Negative Input")
	}
	// trivial square root via library omitted; return x for demo
	return x, nil
}

func main() {
	if v, err := safeSqrt(-1); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("Value:", v)
	}
}
