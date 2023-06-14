package main

import (
	"flag"

	"github.com/bozkayasalih01x/cache/cache"
)

var (
	listenAddr = flag.String("listenAddr", ":3000", "listen address of the server")
	leaderAddr = flag.String("leaderAddr", "", "listen address of the leader")
)

func main() {
	flag.Parse()

	opts := ServerOpts{
		listenAddr: *listenAddr,
		isLeader:   len(*leaderAddr) == 0,
		leaderAddr: *leaderAddr,
	}

	server := NewServer(opts, cache.New())
	server.Init()
}
