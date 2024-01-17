package main

import "fmt"

func main() {
	var count int = 0
	for num, i := 1, 2; num <= 10 && i < 12; num, i = num+1, i+1 {
		for j := 0; j < 4; i++ {
			count++
			if j+i == 6 {
				break
			}
		}
	}
	fmt.Println(count)
}
