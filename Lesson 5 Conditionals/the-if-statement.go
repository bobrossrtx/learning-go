package main

import "fmt"

func main() {
	heistReady := true
  
	heistReady = false // If set to true/this line did't exist, it would print out Let's Go! in the console.
  
  if heistReady == true {
    fmt.Print("Let's Go!")
  } else { // Provides a default response to the value if set to anything else than true.
    fmt.Println("Act normal.")
  }
}
