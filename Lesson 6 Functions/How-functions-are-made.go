package main

import "fmt"

// Define eatTacos() here:
func eatTacos() {
  fmt.Println("Yum!")
}

func main() {
  // Call your function here:
  eatTacos()
}

/*
Above, we defined a function called eatTacos() and, within the body of the function (the part between the curly braces) we print out a message. 
It’s important to note that the code inside the function body does not run until we call the function. 
We call a function by using its name followed by parentheses somewhere outside the definition of the function. 
Our whole main.go file could look like this:
*/

/*
In our example, we defined the function eatTacos() and called it once inside our main() function. 
Notice that our function definition exists outside of main(), but calling eatTacos() occurs inside our main() function. 
This produces the following output in the terminal:

Output:
Yum!
*/
