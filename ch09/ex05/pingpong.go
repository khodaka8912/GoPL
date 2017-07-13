package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan struct{})
	result := make(chan int)
	go echo(ch1, ch2, done, result)
	go echo(ch2, ch1, done, result)
	ch1 <- 0
	<-time.After(1 * time.Second)
	close(done)
	n := <-result
	fmt.Printf("n = %d", n)
}

func echo(in <-chan int, out chan<- int, done <-chan struct{}, result chan<- int) {
	var n int
	for {
		select {
		case n = <-in:
			out <- n + 1
		case <-done:
			result <- n
			return
		}
	}
}
