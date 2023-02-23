package main

import (
	"fmt"

	"github.com/bozkayasalih01x/cache/cache"
	"github.com/bozkayasalih01x/cache/server"
)

func main() {
	opts := server.ServerOptions{
		ListenAddr: ":3000",
		IsLeader:   false,
		LeaderAddr: ":4000",
	}

	fmt.Printf("server started on port %s\n", opts.ListenAddr)

	s := server.New(opts, cache.New())
	s.Start()
}
