package main

import (
	"net"
	"time"
)

func main() {

	con, err := net.Dial("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	_, err = con.Write([]byte("SET maker google"))
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 5)

	_, err = con.Write([]byte("GET maker"))
	if err != nil {
		panic(err)
	}

	defer con.Close()

	select {}

}
