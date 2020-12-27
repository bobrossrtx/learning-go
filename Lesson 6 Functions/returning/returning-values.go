package main

import (
	"fmt"
	"time"
)

// Add "string" as the return type of this function
func isItLateInNewYork() string {
	var lateMessage string
	t := time.Now()
	tz, _ := time.LoadLocation("America/New_York")
	nyHour := t.In(tz).Hour()
	if nyHour < 5 {
		lateMessage = "Goodness it is late"
	} else if nyHour < 16 {
		lateMessage = "It's not late at all!"
	} else if nyHour < 19 {
		lateMessage = "I guess it's getting kind of late"
	} else {
		lateMessage = "It's late"
	}

	// Return the string lateMessage
	return lateMessage
}

func main() {
	var nyLate string
	nyLate = isItLateInNewYork()
	fmt.Println(nyLate)
}

/*
func multiplier(x int32, y int32) int32 {
  return x * y
}

/////////////////////////////////////////
Since both parameters have the same type, we could write it as:

func multiplier(x, y int32) int32 {
  return x * y
}

See how we wrote int32 once at the end of our list of parameters instead of writing int32 after each parameter?

Let’s now call our function with literal values as arguments:
func main() {
  var product int32
  product = multiplier(25, 4)
  fmt.Println(product) // Prints: 100
}

We can also call our function with variables as arguments:
func main() {
  var mainX, mainY, newProduct int32
  mainX = 6
  mainY = 7
  newProduct = multiplier(mainX, mainY)
  fmt.Println(newProduct) // Prints: 42
}
*/
