package main

import "fmt"

func main() {
	n := 3
	switch n {
	case 1:
		fmt.Println("a")
	case 2:
		fmt.Println("b")
	case 3:
		fmt.Println("c")

	default:
		fmt.Println("invalid input")
	}

	a := "a"
	switch a {
	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
	case "c":
		fmt.Println("c")
	case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
		fmt.Println("Number")
	default:
		fmt.Println("invalid input")
	}

	num := 25
	switch {
	case num <= 50:
		if num < 0 {
			break
		}
		fmt.Println("Num less than 50")
		fallthrough
	case num <= 100:
		fmt.Println("Num less than 100")
	case num > 100:
		fmt.Println("Num greater than 100")
	}

}
