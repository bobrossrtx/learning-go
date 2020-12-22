<<<<<<< HEAD
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This now outputs because we set the seed to be random since the system time is always changing, we used the time library.
func main() {
	// Add your code below:
  rand.Seed(time.Now().UnixNano())
  amountLeft := rand.Intn(10000)
  
  fmt.Println("amountLeft is: ", amountLeft)
  
	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
  } else {
    fmt.Println("Where did all my money go?")
  }
}



/* This block of code will always include 81 in the outcome.
func main() {
	// Edit amountLeft below: 
  amountLeft := rand.Intn(10000)
  
  fmt.Println("amountLeft is: ", amountLeft)
  
	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
  } else {
    fmt.Println("Where did all my money go?")
  }
}
*/

/* Topics learned in this lesson:
How to create an if statement that checks a condition and executes code if the condition is true

How to add an else if statement to check for additional conditions.

The use of an else statement that runs by default if none of the previous conditional statements evaluated to true.

Comparison operators check between two operands and evaluate to a boolean.

Logical operators check between two boolean values and return a single boolean.

switch statements can be used to check between multiple conditions much like an if… else if… else statement.

Short variable declarations can be used prior to providing a condition for either if or switch statements. Declared variables are scoped to the statement blocks.

The math/rand library’s .Intn() method is used to generate random numbers.

Go uses a seed value to as a starting point for generating random numbers.

Unique seed values can be created using time, namely rand.seed(time.Now().UnixNano())
*/
=======
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This now outputs because we set the seed to be random since the system time is always changing, we used the time library.
func main() {
	// Add your code below:
  rand.Seed(time.Now().UnixNano())
  amountLeft := rand.Intn(10000)
  
  fmt.Println("amountLeft is: ", amountLeft)
  
	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
  } else {
    fmt.Println("Where did all my money go?")
  }
}



/* This block of code will always include 81 in the outcome.
func main() {
	// Edit amountLeft below: 
  amountLeft := rand.Intn(10000)
  
  fmt.Println("amountLeft is: ", amountLeft)
  
	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
  } else {
    fmt.Println("Where did all my money go?")
  }
}
*/

/* Topics learned in this lesson:
How to create an if statement that checks a condition and executes code if the condition is true

How to add an else if statement to check for additional conditions.

The use of an else statement that runs by default if none of the previous conditional statements evaluated to true.

Comparison operators check between two operands and evaluate to a boolean.

Logical operators check between two boolean values and return a single boolean.

switch statements can be used to check between multiple conditions much like an if… else if… else statement.

Short variable declarations can be used prior to providing a condition for either if or switch statements. Declared variables are scoped to the statement blocks.

The math/rand library’s .Intn() method is used to generate random numbers.

Go uses a seed value to as a starting point for generating random numbers.

Unique seed values can be created using time, namely rand.seed(time.Now().UnixNano())
*/
>>>>>>> ed9344e4fa57a3a383756291d5d0b20d457815be
