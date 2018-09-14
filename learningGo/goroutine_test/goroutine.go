package main

import (
	"fmt"
)

func it1(from string) {
	i := 0
	for j := 0; j < 100; j++ {
		i++
		fmt.Println(from, i)
	}
}

func it2(from string) {
	i := 0
	for j := 0; j < 100; j++ {
		i++
		fmt.Println(from, i)
	}
}

func main() {
	go it1("func1")

	go it2("func2")

	fmt.Scanln()
	fmt.Println("Done")
}
