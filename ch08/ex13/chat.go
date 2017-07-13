package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	ch   chan<- string
	name string
}

var (
	entering = make(chan *client)
	leaving  = make(chan *client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[*client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			members := bytes.NewBufferString("members: ")
			first := true
			for other := range clients {
				if first {
					first = false
				} else {
					members.WriteString(", ")
				}
				members.WriteString(other.name)
			}
			cli.ch <- members.String()
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := client{ch, conn.RemoteAddr().String()}
	ch <- "You are " + who.name
	messages <- who.name + " has arrived"

	entering <- &who

	input := bufio.NewScanner(conn)
	inputChan := make(chan string)
	go func() {
		for {
			select {
			case msg := <-inputChan:
				messages <- msg
			case <-time.After(5 * time.Minute):
				ch <- "disconnect silent user"
				conn.Close()
				return
			}
		}
	}()

	for input.Scan() {
		inputChan <- who.name + ": " + input.Text()
	}

	leaving <- &who
	messages <- who.name + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
