package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/bozkayasalih01x/cache/cache"
)

type ServerOpts struct {
	listenAddr string
	leaderAddr string
	isLeader   bool
}

type Server struct {
	opts    ServerOpts
	cache   cache.Cacher
	clients map[net.Conn]struct{}
}

func NewServer(opts ServerOpts, c cache.Cacher) *Server {
	return &Server{
		opts:    opts,
		clients: make(map[net.Conn]struct{}),
		cache:   c,
	}
}

func (s *Server) Init() {
	listener, err := net.Listen("tcp", *listenAddr)

	fmt.Printf("cache started on %s\n", *listenAddr)
	if err != nil {
		log.Fatalf("cound't create new server: %v", err)
	}


	// NOTE: handle dial when its not leader
	if !s.opts.isLeader {
		go func() {
			conn, err := net.Dial("tcp", *listenAddr)
			if err != nil {
				log.Printf("couldn't dial to conn: %v", err)
				return
			}
			s.handleConn(conn)
		}()
	}

  
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("couldn't connect to net: %v", err)
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()
  
	fmt.Printf("connection made %s to %s\n", conn.LocalAddr().String(), conn.RemoteAddr().String())
	buf := make([]byte, 1024)

	if s.opts.isLeader {
		s.clients[conn] = struct{}{}
	}


	buf := make([]byte, 1024)
  
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Printf("couldn't read the bytes: %v\n", err)
			break
		}
		go s.handleMsg(conn, buf[:n])
	}
}

type Cmd string

const (
	CmdSet  Cmd = "SET"
	CmdGet  Cmd = "GET"
	CmdJoin Cmd = "JOIN"
)

func (s *Server) handleMsg(conn net.Conn, rawData []byte) {
	msg, err := s.parseMessage(rawData)
	if err != nil {
		fmt.Printf("failed to parse command: %v\n", err)
		conn.Write([]byte(err.Error()))
		return
	}
	fmt.Printf("msg received: %v\n", msg)

	switch msg.Cmd {
	case CmdSet:
		err = s.handleSetCmd(conn, msg)
	case CmdGet:
		err = s.handleGetCmd(conn, msg)
	case CmdJoin:
		err = s.handleJoinCmd(conn, msg)
	default:
		log.Printf("invalid cmd \n")
		err = errors.New("invalid cmd detected\n")
	}
	if err != nil {
		fmt.Printf("errors: %v\n", err)
		conn.Write([]byte(err.Error()))
	}
}

type Message struct {
	Cmd   Cmd
	Key   string
	Value string
	TTL   time.Duration
}

func (s *Server) parseMessage(rawData []byte) (*Message, error) {
	msg := &Message{}

	strSlice := strings.Split(string(rawData), " ")

	if len(strSlice) < 2 {
		log.Printf("invalid cmd")
		return nil, errors.New("invalid cmd pased, check cmd again!")
	}

	msg.Cmd = Cmd(strSlice[0])
	msg.Key = strSlice[1]

	if len(strSlice) > 2 {
		ttl, err := strconv.Atoi(strSlice[3])
		if err != nil {
			return nil, err
		}
		msg.Value = strSlice[2]
		msg.TTL = time.Duration(ttl)
	}
	return msg, nil
}

func (s *Server) handleSetCmd(conn net.Conn, msg *Message) error {
	err := s.cache.Set([]byte(msg.Key), []byte(msg.Value), time.Duration(msg.TTL))
	if err != nil {
		return err
	}

	go s.sendToMembers(context.TODO(), msg)
	return nil
}

func (s *Server) handleGetCmd(conn net.Conn, msg *Message) error {
	resp, err := s.cache.Get([]byte(msg.Key))
	if err != nil {
		return errors.New("not found such a data\n")
	}

	// NOTE: print out the resp test demostration
	if _, err := conn.Write(resp); err != nil {
		return fmt.Errorf("couldn't write to connection: %v\n", err)
	}
	fmt.Printf("msg sending back to %s\n", conn.LocalAddr().String())
	_, err = conn.Write(resp)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) handleJoinCmd(conn net.Conn, msg *Message) error {
	// NOTE: not implemented yet. working on
	return nil
}

func (s *Server) sendToMembers(ctx context.Context, msg *Message) error {
	for client := range s.clients {
		rawMsg := msg.ToBytes()
		fmt.Printf("distributing to %s\n", client.LocalAddr().String())
		fmt.Printf("forwarding raw msg to followers %s\n", string(rawMsg))
		_, err := client.Write(rawMsg)
		if err != nil {
			fmt.Printf("couldn't write to follower %v\n", err)
			continue
		}
	}
	return nil
}
