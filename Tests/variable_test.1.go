package main

import "fmt"

func main() {

	var publisher, writer, artist, title, genres string

	var year, pageNumber, firstCost, secondCost, thirdCost, age int

	var grade float32

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	age = 2020 - 1997
	pageNumber = 14
	grade = 6.7
	firstCost = 7 * (pageNumber / age)

	fmt.Println(title, ":", "written by", writer)
	fmt.Println("drawn by", artist)
	fmt.Println("Release Date :", year)
	fmt.Println("publisher :", publisher)
	fmt.Println("Number of pages :", pageNumber, "Grade :", grade)
	fmt.Println("Cost :", "$", firstCost)

	fmt.Println()

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	year = 2013
	age = 2020 - 2013
	pageNumber = 160
	grade = 9.0
	secondCost = 9 * ((pageNumber / 30) / age)

	fmt.Println(title, ":", "written by", writer)
	fmt.Println("drawn by", artist)
	fmt.Println("Release Date :", year)
	fmt.Println("publisher :", publisher)
	fmt.Println("Number of pages :", pageNumber, "Grade :", grade)
	fmt.Println("Cost :", "$", secondCost)

	fmt.Println()

	title = "Batman: The Killing Joke"
	writer = "Alan Moore"
	artist = "Brian Bolland"
	year = 1988
	age = 32
	pageNumber = 48
	grade = 9.6
	genres = "Comics, Comic book, Novel, Graphic novel, Fantasy Fiction"
	thirdCost = 10 * (pageNumber / age)
	// thirdCost is the pagenumber tho then be divided by th age which is then multiplied by it's grade which is rounded to its closest hole number.

	fmt.Println(title, ":", "written by", writer)
	fmt.Println("drawn by", artist)
	fmt.Println("Release Date :", year)
	fmt.Println("publisher :", publisher)
	fmt.Println("Genres :", genres)
	fmt.Println("Number of pages :", pageNumber, "Grade :", grade)
	fmt.Println("Cost :", "$", thirdCost)

}
