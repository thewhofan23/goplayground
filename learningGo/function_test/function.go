package main

import (
	"fmt"
	"strconv"
)

func forloop() string {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	return strconv.Itoa(sum)
}

func arrayloop(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}

func main() {
	slice := []int{1, 2, 3}
	output := forloop()
	fmt.Println(output)
	arrout := arrayloop(slice)
	fmt.Println(arrout)
}
