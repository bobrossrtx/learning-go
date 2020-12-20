package main

import (
	"fmt"
)

func main() {
	var count = 0
	var projects = 3 * 2

	count += projects
	var learnGo, calculator, onTopServers, portfolio, application, propaganda string
	learnGo = "LearnGo : Decided to learn Go coding language."
	calculator = "Calculator : Made a simple calculator using JavaScript."
	onTopServers = "OnTopServers : Discord server promotion site (not public)"
	portfolio = "Portfolio : the portfolio of Owen Boreham (bobrossrtx)."
	application = "Android random word generator : A simple random word generator in the form of an android application"
	propaganda = "Propaganda : Propaganda is a open source css framework free to use by anyone building a webpage"

	fmt.Println("Projects")
	fmt.Println()
	fmt.Println("Count :", count)
	fmt.Println()
	fmt.Println("Titles :")
	fmt.Println(learnGo)
	fmt.Println(calculator)
	fmt.Println(onTopServers)
	fmt.Println(portfolio)
	fmt.Println(application)
	fmt.Println(propaganda)
}
