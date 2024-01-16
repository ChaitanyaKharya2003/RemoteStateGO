package main

import "fmt"
import "math"

func add(num1, num2 int) int {
	sum := num1 + num2
	return sum
}

func RectangleProps(l, b float64) (float64, float64) {
	area := l * b
	peri := 2 * (l + b)
	return area, peri
}

func CircleProps(radius float64) (circumference, area float64) {
	area = math.Pi * math.Pow(radius, 2)
	circumference = 2 * math.Pi * radius
	return
}

func main() {
	a, b := 5, 6
	sum := add(a, b)
	fmt.Println(sum)

	var length, breadth float64 = float64(10), 12.356
	area, peri := RectangleProps(length, breadth)
	fmt.Println("area:", area, "perimeter:", peri)

	rad := 5.49876
	var circum, areaC float64 = CircleProps(rad)
	fmt.Printf("Circumference of circle with radius %.2f is %.2f and area is %.2f\n", rad, circum, areaC)

	// _ is blank identifier
	_, periRect := RectangleProps(4.56, 6.53)
	fmt.Println("Perimeter:", periRect)
}
