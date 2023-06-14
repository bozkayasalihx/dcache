package main

import (
	"fmt"
	"log"
)

func (m *Message) ToBytes() []byte {
	var cmd string
	switch m.Cmd {
	case CmdSet:
		cmd = fmt.Sprintf("%s %s %s %d", m.Cmd, m.Key, m.Value, m.TTL)
		return []byte(cmd)
	case CmdGet:
		cmd = fmt.Sprintf("%s %s", m.Cmd, m.Key)
		return []byte(cmd)
	case CmdJoin:
		cmd = fmt.Sprintf("%s %s", m.Cmd, m.Key)
		return []byte(cmd)
	default:
		log.Printf("invalid cmd")
	}

	return nil
}
