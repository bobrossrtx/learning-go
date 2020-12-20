package main

// What the package name wll be when you package you go file to an exe

// Import packages with import

import (
	"fmt"
	t "time"
)

// You can make an alias the package my writing a name befor the package
// For example, import t "time"

func main() {
	fmt.Println("Time: ", t.Now())
}
