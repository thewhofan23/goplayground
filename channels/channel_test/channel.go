package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- false
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}

// func main() {
// 	message := make(chan string)

// 	go func() {
// 		message <- "test"
// 	}()

// 	fmt.Println(<-message)
// }

// func main() {
// 	channel := make(chan string, 2)

// 	channel <- "1"
// 	channel <- "2"

// 	fmt.Println(<-channel)
// 	fmt.Println(<-channel)
// }
