package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	wg := sync.WaitGroup{}
	for input.Scan() {
		wg.Add(1)
		go echo(c, input.Text(), 1*time.Second, wg)
	}
	switch c.(type) {
	case *net.TCPConn:
		c.(*net.TCPConn).CloseWrite()
		wg.Wait()
		c.(*net.TCPConn).CloseRead()
	default:
		log.Fatalf("unexpected connection:%t", c)
		c.Close()
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
