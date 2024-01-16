package main

import "fmt"

func main() {
	var age int = 20
	var name string = "Chaitanya"
	var first, last string = "Chaitanya", "Kharya"
	var year = 3
	var hello string
	fmt.Println("My Name is", name, "and my age is", age)
	hello = "Hello World"
	fmt.Println("Full Name:", first, last)
	fmt.Println("Year:", year)
	fmt.Println(hello)

	var (
		lang             = "Go"
		assignmentNumber = 3
		completed        = false
	)
	fmt.Println(lang, assignmentNumber, completed)

	date := "01/01"
	fmt.Println(date)
	completed = true
	fmt.Println("Completed:", completed)
}
