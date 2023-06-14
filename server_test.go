package main

import (
	"context"
	"net"
	"reflect"
	"testing"

	"github.com/bozkayasalih01x/cache/cache"
)

func TestNewServer(t *testing.T) {
	type args struct {
		opts ServerOpts
		c    cache.Cacher
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.opts, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Init(t *testing.T) {
	type fields struct {
		opts    ServerOpts
		cache   cache.Cacher
		clients map[net.Conn]struct{}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				opts:    tt.fields.opts,
				cache:   tt.fields.cache,
				clients: tt.fields.clients,
			}
			s.Init()
		})
	}
}

func TestServer_handleConn(t *testing.T) {
	type fields struct {
		opts    ServerOpts
		cache   cache.Cacher
		clients map[net.Conn]struct{}
	}
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				opts:    tt.fields.opts,
				cache:   tt.fields.cache,
				clients: tt.fields.clients,
			}
			s.handleConn(tt.args.conn)
		})
	}
}

func TestServer_handleMsg(t *testing.T) {
	type fields struct {
		opts    ServerOpts
		cache   cache.Cacher
		clients map[net.Conn]struct{}
	}
	type args struct {
		conn    net.Conn
		rawData []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				opts:    tt.fields.opts,
				cache:   tt.fields.cache,
				clients: tt.fields.clients,
			}
			if err := s.handleMsg(tt.args.conn, tt.args.rawData); (err != nil) != tt.wantErr {
				t.Errorf("Server.handleMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_parseMessage(t *testing.T) {
	type fields struct {
		opts    ServerOpts
		cache   cache.Cacher
		clients map[net.Conn]struct{}
	}
	type args struct {
		rawData []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				opts:    tt.fields.opts,
				cache:   tt.fields.cache,
				clients: tt.fields.clients,
			}
			got, err := s.parseMessage(tt.args.rawData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.parseMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.parseMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_handleSetCmd(t *testing.T) {
	type fields struct {
		opts    ServerOpts
		cache   cache.Cacher
		clients map[net.Conn]struct{}
	}
	type args struct {
		conn net.Conn
		msg  *Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				opts:    tt.fields.opts,
				cache:   tt.fields.cache,
				clients: tt.fields.clients,
			}
			if err := s.handleSetCmd(tt.args.conn, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Server.handleSetCmd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_handleGetCmd(t *testing.T) {
	type fields struct {
		opts    ServerOpts
		cache   cache.Cacher
		clients map[net.Conn]struct{}
	}
	type args struct {
		conn net.Conn
		msg  *Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				opts:    tt.fields.opts,
				cache:   tt.fields.cache,
				clients: tt.fields.clients,
			}
			if err := s.handleGetCmd(tt.args.conn, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Server.handleGetCmd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_handleJoinCmd(t *testing.T) {
	type fields struct {
		opts    ServerOpts
		cache   cache.Cacher
		clients map[net.Conn]struct{}
	}
	type args struct {
		conn net.Conn
		msg  *Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				opts:    tt.fields.opts,
				cache:   tt.fields.cache,
				clients: tt.fields.clients,
			}
			if err := s.handleJoinCmd(tt.args.conn, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Server.handleJoinCmd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_sendToMembers(t *testing.T) {
	type fields struct {
		opts    ServerOpts
		cache   cache.Cacher
		clients map[net.Conn]struct{}
	}
	type args struct {
		ctx context.Context
		msg *Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				opts:    tt.fields.opts,
				cache:   tt.fields.cache,
				clients: tt.fields.clients,
			}
			if err := s.sendToMembers(tt.args.ctx, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Server.sendToMembers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
