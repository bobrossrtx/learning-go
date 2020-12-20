package main

import "fmt"

func main() {

	// Not operator
	
	readyToGo := true

	if !readyToGo {
		fmt.Println("Start the car!")
	} else {
		fmt.Println("What are we waiting for??")
	}
}
