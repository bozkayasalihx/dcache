package main

import (
	"flag"
	"fmt"

	"github.com/bozkayasalih01x/cache/cache"
	"github.com/bozkayasalih01x/cache/server"
)

func main() {
	var listenAddr string
	var leaderAddr string
	flag.StringVar(&listenAddr, "a", "", "set up listen address")
	flag.StringVar(&leaderAddr, "la", "", "set up leader address")
	flag.Parse()

	if listenAddr == "" || leaderAddr == "" {
		fmt.Printf("%s or %s must be setted", listenAddr, leaderAddr)
		flag.Usage()
		return
	}

	opts := server.ServerOptions{
		ListenAddr: listenAddr,
		IsLeader:   false,
		LeaderAddr: leaderAddr,
	}

	fmt.Printf("server started on port %s\n", opts.ListenAddr)

	s := server.New(opts, cache.New())
	s.Start()
}
