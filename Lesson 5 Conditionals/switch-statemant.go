package main

import "fmt"

func main() {
	name := "H. J. Simp"

	// Add your switch statement below:
  switch name {
    case "butch":
      fmt.Println("Head to Robbers Roost.")
    case "Bonnie":
      fmt.Println("Stay put in joplin")
    default:
      fmt.Println("Just hide!")
  }
}

// Output: Just Hide!