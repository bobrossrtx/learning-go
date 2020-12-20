package main

import "fmt"

func main() {
	vaultAmt := 2356468
	
  // Add your code below:
  var bag string
  var enough string
  bag = "We're going to need more bags!"
  enough = "Were got enough bags, come on!"

  if vaultAmt >= 200000 {
    fmt.Printf("%v", bag)
  } else {
    fmt.Printf("%v", enough)
  }
  
}
