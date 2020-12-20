package main

import "fmt"

func main() {
	// Add a fmt.Println() statement
	// that prints 2235 * 1231

	fmt.Println("2235 * 1231 = ", 2235*1231)
	fmt.Println()

	// Create the constant earthsGravity
	// here and set it to 9.80665
	const earthsGravity = 9.80665

	// Here's where we print out the gravity:
	fmt.Println("Earths gravity :", earthsGravity)
	fmt.Println()

	const funFact = "Hummingbirds' wings can beat up to 200 times a second."

	fmt.Println("FunFact")
	fmt.Println("Did you know? :", funFact)
	fmt.Println()

	var numOfFlavors int
	// Assign a value for numOfFlavors below:
	numOfFlavors = 57

	fmt.Println("Flavor count :", numOfFlavors)

	// Declare flavorScale below:
	var flavorScale float32 = 5.8

	fmt.Println("Flavor Scale", flavorScale)
	fmt.Println()

	// Define cupsOfCoffeeConsumed here
	var cupsOfCoffeeConsumed int

	// Give a value to cupsOfCoffeeConsumed
	cupsOfCoffeeConsumed = 2

	// Print out cupsOfCoffeeConsumed
	fmt.Println("I drank", cupsOfCoffeeConsumed, "coffees")
	fmt.Println()

	// Define magicNum and powerLevel below:
	var magicNum, powerLevel int32
	magicNum = 2048
	powerLevel = 9001

	fmt.Println("magicNum is:", magicNum, "powerLevel is:", powerLevel)

	// Define amount and unit below:
	amount, unit := 10, "doll hairs"

	fmt.Println(amount, unit, ", that's expensive...")
}

/*
import "fmt"

func main() {
  var stationName string
  var nextTrainTime int8
  var isUptownTrain bool

  stationName = "Union Square"
  nextTrainTime = 12
  isUptownTrain = false

  fmt.Println("Current station:", stationName)
  fmt.Println("Next train:", nextTrainTime, "minutes")
  fmt.Println("Is uptown:", isUptownTrain)

  stationName = "Grand Central"
  nextTrainTime = 3
  isUptownTrain = true

  fmt.Println("")
  fmt.Println("Current station:", stationName)
  fmt.Println("Next train:", nextTrainTime, "minutes")
  fmt.Println("Is uptown:", isUptownTrain)
}
*/
