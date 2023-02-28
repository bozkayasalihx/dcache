package server

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/bozkayasalih01x/cache/cache"
	proto "github.com/bozkayasalih01x/cache/protocols"
)

type Message struct {
	Key   []byte
	Value []byte
}

type ServerOptions struct {
	ListenAddr string
	IsLeader   bool
	LeaderAddr string
}

type Server struct {
	ServerOptions
	c cache.Cacher
}

func New(opts ServerOptions, c cache.Cacher) *Server {
	return &Server{
		c:             c,
		ServerOptions: opts,
	}
}

func (s *Server) Start() error {
	lns, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return fmt.Errorf("couldn't listen the network %v", err)
	}

	for {
		conn, err := lns.Accept()
		if err != nil {
			if err == io.EOF {
				fmt.Println("no more data to read")
				break
			}
			fmt.Printf("couldn't accept the connection %v", err)
			continue
		}
		go s.handleConnection(conn)
	}
	return nil
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("connection made with %s\n", conn.RemoteAddr())
	for {
		cmd, err := proto.ParseCommand(conn)
		if err != nil {
			fmt.Printf("an error accured try again later!\n %v", err)
			break
		}
		go s.handleCommand(conn, cmd)

	}
	fmt.Printf("connection closed with %s", conn.RemoteAddr())
}

func (s *Server) handleCommand(conn net.Conn, cmd interface{}) {
	switch v := cmd.(type) {
	case *proto.MessageGetType:
		s.handleGetCommand(conn, v)
	case *proto.MessageSetType:
		s.handleSetCommand(conn, v)
	}
}

func (s *Server) handleSetCommand(conn net.Conn, cmd *proto.MessageSetType) {
	fmt.Printf("key is  -> %s and value is %s \n", cmd.Key, cmd.Value)
	err := s.c.Set(cmd.Key, cmd.Value)
	if err != nil {
		log.Fatalf("couldnt set %s to %s storage", string(cmd.Key), string(cmd.Value))

	}
}

func (s *Server) handleGetCommand(conn net.Conn, cmd *proto.MessageGetType) {
	val, err := s.c.Get(cmd.Key)
	if err != nil {
		log.Fatalf("couldnt get the data from cache %v  ", err)
	}

	fmt.Printf("the data from %s connection is  -> %s \n", conn.RemoteAddr(), string(val))
	_, err = conn.Write(val)
	if err != nil {
		panic(err)
	}
}
