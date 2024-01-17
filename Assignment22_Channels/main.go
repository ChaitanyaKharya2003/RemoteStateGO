package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendData(ch chan<- int) {
	ch <- 12
}

func hello(done chan bool, data int) {
	fmt.Println("Hello")
	fmt.Println("Goroutine before sleep")
	time.Sleep(time.Second * 2)
	fmt.Println("Goroutine Awake")
	if data%2 == 0 {
		done <- true
	} else {
		done <- false
	}
}

func assign(a chan int) {
	num := rand.Int()
	a <- num
}
func main() {
	var a chan int
	if a == nil {
		fmt.Println("Channel is nil, defining")
		a = make(chan int)
		fmt.Printf("Type of channel a is %T\n", a)
	}

	go assign(a)
	data := <-a
	fmt.Println(data, "from channel")

	done := make(chan bool)

	go hello(done, data)
	booldata := <-done
	if booldata {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	unichannel := make(chan int)
	go sendData(unichannel)
	fmt.Println(<-unichannel)
}
