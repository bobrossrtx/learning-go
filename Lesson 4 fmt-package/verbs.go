package main

import "fmt"

func main() {

  // printf / verbs
  // The verbs: %v, %T, %d, and %f.
  floatExample := 1.75
  // Edit the following Printf for the FIRST step
  fmt.Printf("Working with a %T", floatExample) 
  
  fmt.Println("\n***") // Added for spacing
  
  yearsOfExp := 3
  reqYearsExp := 15
  // Edit the following Printf for the SECOND step
  fmt.Printf("I have %d years of Go experience and this job is asking for %d years", yearsOfExp, reqYearsExp) 
  
  fmt.Println("\n***") // Added for spacing
  
  stockPrice := 3.50
  // Edit the following Printf for the THIRD step
  fmt.Printf("Each share of gopher feed is $%.2f!", stockPrice) 

  fmt.Println("\n***") // Added for spacing

  // sprint f
  template := "I wish I had a %v."
  pet := "puppy"
  
  var wish string
  
  // Add your code below:
  wish = fmt.Sprintf( template, pet )
  
  fmt.Println(wish)

  // fmt.Scan
  // Used for collecting user responses
  fmt.Println("What would you like for lunch?")
  fmt.Println("(choose 2 foods. ie: pizza chips)")
  
  // Add your code below:
  var food1 string
  var food2 string
  fmt.Scan(&food1)
  fmt.Scan(&food2)
  
  fmt.Printf("Sure, we can have %v and %v for lunch.", food1, food2)
}






