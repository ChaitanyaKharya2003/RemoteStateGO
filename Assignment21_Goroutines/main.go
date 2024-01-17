package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(250 * time.Millisecond)

	}
}

func alpha() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(400 * time.Millisecond)

	}
}

func hello() {
	fmt.Println("Hello World Goroutines")
}
func main() {
	go hello()
	fmt.Println("Hello from main")
	time.Sleep(1 * time.Second)

	go numbers()
	go alpha()
	time.Sleep(3 * time.Second)
	fmt.Println("Main terminated")
}
