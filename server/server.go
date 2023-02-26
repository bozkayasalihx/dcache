package server

import (
	"fmt"
	"io"
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
			fmt.Printf("an error accured try again later! %v", err)
			break
		}
		go s.handleCommand(cmd)

	}
	fmt.Printf("connection closed with %s", conn.RemoteAddr())
}

func (s *Server) handleCommand(cmd interface{}) {
	switch cmd.(type) {
	case *proto.MessageGetType:
	case *proto.MessageSetType:
	default:
		panic("command not found")
	}
}

func (s *Server) handleSetCommand(cmd *proto.MessageSetType) {
	fmt.Printf("the set command is this %v", cmd)
}

func (s *Server) handleGetCommand(cmd *proto.MessageGetType) {
	fmt.Printf("the get command is this %v", cmd)
}
