package main

import (	
	f "fmt"
	t "time"
)

func main() {

	var time = t.Now()

	f.Print(time)
}

// output: 2020-12-19 16:10:54.8376756 +0000 GMT m=+0.006988201
// ToDo challenge: to be able to shorten down to just the date, and to convert Into an integer.