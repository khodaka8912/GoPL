package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	clocks := map[string]string{}
	for _, arg := range os.Args[1:] {
		clock := strings.Split(arg, "=")
		if len(clock) != 2 {
			log.Fatalf("invalid clock: %q", arg)
			continue
		}
		clocks[clock[0]] = clock[1]
	}
	fmt.Printf("clocks:\n%v\n", clocks)
	length := len(clocks)
	ch := make([]chan string, length)
	i := 0
	for name, address := range clocks {
		ch[i] = make(chan string)
		go readClock(name, address, ch[i])
		i++
	}
	times := make([]string, length)
	ok := true
	for {
		for i, ch := range ch {
			times[i], ok = <-ch
			if !ok {
				return
			}
		}
		fmt.Println(strings.Join(times, "\t"))
	}
}

func readClock(name, address string, ch chan<- string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	defer close(ch)
	ch <- fmt.Sprintf("%8s", name)
	reader := bufio.NewReader(conn)
	for bytes, _, err := reader.ReadLine(); err == nil; bytes, _, err = reader.ReadLine() {
		ch <- string(bytes)
	}
}
