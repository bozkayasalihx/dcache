package server

import (
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/bozkayasalih01x/cache/cache"
)

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
	cmd := make([]byte, 2048)
	for {
		n, err := conn.Read(cmd)
		if err != nil {
			fmt.Printf("an error accured try again later! %v", err)
			break
		}
		go s.handleCommand(cmd[:n])

	}
	fmt.Printf("connection closed with %s", conn.RemoteAddr())
}

type MessageSet struct {
	Key   []byte
	Value []byte
}

type MessageGet struct {
	Key []byte
}

var (
	MSGSet = "SET"
	MSGGet = "GET"
)

func (s *Server) handleCommand(cmd []byte) error {
	cmdArray := strings.Split(string(cmd), " ")
	cmdHeader := cmdArray[0]
	var err error

	switch cmdHeader {
	case MSGGet:
		err = s.commandGet(cmdArray)
	case MSGSet:
		err = s.commandSet(cmdArray)
	default:
		err = fmt.Errorf("couldn't identify command %v", err)
	}
	return err
}

func (s *Server) commandSet(cmdArray []string) error {
	cmdSet := &MessageSet{}
	if len(cmdArray) > 3 {
		return fmt.Errorf("couldn't longer than 3")
	}
	cmdSet.Key = []byte(cmdArray[1])
	cmdSet.Value = []byte(cmdArray[2])

	s.handleSetCommand(cmdSet)
	return nil
}

func (s *Server) commandGet(cmdArray []string) error {
	cmdGet := &MessageGet{}
	if len(cmdArray) > 2 {
		return fmt.Errorf("command couldn't longer than 2")
	}
	cmdGet.Key = []byte(cmdArray[1])

	s.handleGetCommand(cmdGet)
	return nil
}

func (s *Server) handleGetCommand(cmdGet *MessageGet) {
	data, err := s.c.Get(cmdGet.Key)
	if err != nil {
		fmt.Errorf("couldn't get %s from store", string(cmdGet.Key))
	}

	fmt.Printf("the data that you stored is %s\n", string(data))
}

func (s *Server) handleSetCommand(cmdSet *MessageSet) {
	err := s.c.Set(cmdSet.Key, cmdSet.Value)
	if err != nil {
		fmt.Errorf("couldn't set %s -> %s to store", string(cmdSet.Key), string(cmdSet.Value))
	}
}
