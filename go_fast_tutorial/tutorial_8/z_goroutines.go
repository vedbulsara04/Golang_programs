package main

import (
  "fmt"
  "time"
)

func printMessage(msg string){
  for i := 0; i < 5; i++ {
    fmt.Println(msg, i)
    time.Sleep(500*timee.Milisecond)
  }
}

func main() {
  // Start a goroutine
  go printMessage("Hello from Goroutine!")

  // Run in main goroutine
  printMessage("Hello from Main")
}
