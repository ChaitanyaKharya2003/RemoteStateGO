package main

import "fmt"

func main() {
	if num := 99; num < 50 {
		fmt.Println("Num less than 50")
		return
	} else if num >= 51 && num < 100 {
		fmt.Println("Num between 51 and 100")
		return
	}
	fmt.Println("Num greater than 100")
}
