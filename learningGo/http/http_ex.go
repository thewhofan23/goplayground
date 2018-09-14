package main

import (
	"io"
	"net/http"
	"strconv"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	i := make(chan int)
	temp := 0
	go func() {
		for {
			i <- temp
			time.Sleep(time.Second)
			temp++
		}
	}()
	for j := 0; j < 5; j++ {
		io.WriteString(w, strconv.Itoa(<-i))
	}
	close(i)

}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		panic(err)
	}
}
