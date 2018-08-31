package main

import "fmt"

func ping(pings chan string, msg string) {
	pings <- msg
}

func pong(pings chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string)
	pongs := make(chan string, 1)
	go ping(pings, "testing this shit")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
