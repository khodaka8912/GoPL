package main

import (
	"flag"
	"fmt"
	"time"
)

var limit = flag.Int("limit", 0, "limit of goroutin to run.")

func main() {
	flag.Parse()
	start := make(chan struct{})
	in := start
	for i := 0; *limit <= 0 || i < *limit; i++ {
		out := make(chan struct{})
		go pipe(in, out)
		in = out
	}
	end := in
	now := time.Now()
	start <- struct{}{}
	<-end
	elapsed := time.Now().Sub(now)
	fmt.Printf("%f ms has elapsed for communications of %d pipes", elapsed.Seconds()*1000, *limit)

}

func pipe(in <-chan struct{}, out chan<- struct{}) {
	for data := range in {
		out <- data
	}
	close(out)
}
