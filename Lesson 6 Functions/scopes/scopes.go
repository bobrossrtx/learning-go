package main

import "fmt"


/* This code will not work since the instructions are undefined within the 
   main function but it is in the startGame() function 

func startGame() {
  instructions := "Press enter to start..." 
  
  
}

func main() {
  startGame()
  
  fmt.Println(instructions)
}
*/

// Here is the fix

func startGame() {
	instructions := "Press enter to start..." 
	
	fmt.Println(instructions)
  }
  
  func main() {
	startGame()
	
  }