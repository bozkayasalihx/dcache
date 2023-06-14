package main

import (
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Dial("tcp", ":4000")
	defer l.Close()
	if err != nil {
		log.Fatalf("couldn't connect to net: %v", err)
	}
	l.Write([]byte("SET google fuck_the_ping 0"))

	time.Sleep(2 * time.Second)

	l.Write([]byte("GET google"))
}
