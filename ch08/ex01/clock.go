package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var port = flag.Int("port", 8000, "port num to listen")
var tz = flag.String("tz", "Asia/Tokyo", "time zone of clock")

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(locale).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

var locale *time.Location

func main() {
	flag.Parse()
	if *port < 0 || *port > 65535 {
		log.Fatalf("invalid port num: %d", *port)
		os.Exit(1)
	}
	loc, err := time.LoadLocation(*tz)
	if err != nil {
		log.Fatalf("invalid location: %s", *tz)
		locale = time.Local
	} else {
		locale = loc
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}
