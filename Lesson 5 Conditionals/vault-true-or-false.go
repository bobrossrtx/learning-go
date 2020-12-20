package main

import "fmt"

func main() {

	// True or false operator/Is equal to
	
	lockCombo := "21-35-9"
	robAttempt := "1-1-1"

  // Add your code below:
  if (lockCombo == robAttempt) {
    fmt.Println("The vault is now opened.")
  } else {
    fmt.Println("Police have been called.")
  }
  
}
