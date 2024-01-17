package main

import "fmt"

func sumSquares(num int, squarech chan int) {
	var sum int = 0
	for num != 0 {
		num1 := num % 10
		sum += num1 * num1
		num /= 10
	}
	squarech <- sum
}

func sumCubes(num int, cubech chan int) {
	var sum int = 0
	for num != 0 {
		num1 := num % 10
		sum += num1 * num1 * num1
		num /= 10
	}
	cubech <- sum
}

func main() {
	squarech := make(chan int)
	cubech := make(chan int)
	go sumSquares(5678, squarech)
	go sumCubes(5678, cubech)
	squares, cubes := <-squarech, <-cubech
	fmt.Println(squares + cubes)
}
