package main

import (
	"fmt"
	"strings"
)

func main() {
	url := "Mr John Smith    "
	length := 13
	spaceCount := 0
	index := 0
	for i := 0; i < length; i++ {
		if url[i] == ' ' {
			spaceCount++
		}
	}
	index = length + 2*spaceCount
	urls := strings.Split(url, "")
	for j := length - 1; j >= 0; j-- {
		if urls[j] == " " {
			urls[index-3] = "%"
			urls[index-2] = "2"
			urls[index-1] = "0"
			index = index - 3
		} else {
			urls[index-1] = urls[j]
			index--
		}
	}
	fmt.Println(strings.Join(urls, ""))
}
