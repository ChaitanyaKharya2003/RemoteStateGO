package main

import "fmt"

func add(num1, num2 int) int {
	sum := num1 + num2
	return sum
}

func RectangleProps(l, b float64) (float64, float64) {
	area := l * b
	peri := 2 * (l + b)
	return area, peri
}

func main() {
	a, b := 5, 6
	sum := add(a, b)
	fmt.Println(sum)
}