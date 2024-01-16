package main

import "fmt"

func main() {
	const a int = 10
	fmt.Println(a)
	const (
		name    = "Chaitanya"
		age     = 20
		country = "India"
	)
	var n = name
	fmt.Printf("type: %T, value: %v\n", n, n)

	var b = 5.9 / 9
	fmt.Printf("type: %T, value: %v", b, b)

}
