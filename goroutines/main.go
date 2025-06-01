package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello", phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	// dones := make([]chan bool, 3)

	done := make(chan bool)

	// dones[0] = make(chan bool)
	// dones[1] = make(chan bool)
	// dones[2] = make(chan bool)

	go greet("1st fast", done)
	go slowGreet("slow world", done)
	go greet("2nd fast", done)

	for range done {
	}
}
