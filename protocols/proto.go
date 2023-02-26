package proto

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Command byte

const (
	MSGSET Command = iota
	MSGGET
	MSGJOIN
)

type MessageGetType struct {
	Key []byte
}

type MessageSetType struct {
	Key   []byte
	Value []byte
}

func ParseCommand(r io.Reader) (interface{}, error) {
	var command Command
	err := binary.Read(r, binary.LittleEndian, &command)
	if err != nil {
		return fmt.Errorf("couldnt read the command %v", err), nil
	}

	switch command {
	case MSGGET:
		return handleGetCommand(r)
	case MSGSET:
		return handleSetCommand(r)
	default:
		return nil, fmt.Errorf("couldnt identify the command\n")
	}
}

func handleGetCommand(r io.Reader) (*MessageGetType, error) {
	cmd := &MessageGetType{}
	var keyLen int32
	err := binary.Read(r, binary.LittleEndian, &keyLen)
	if err != nil {
		return nil, err
	}
	cmd.Key = make([]byte, keyLen)
	err = binary.Read(r, binary.LittleEndian, &cmd.Key)
	return cmd, err

}

func handleSetCommand(r io.Reader) (*MessageSetType, error) {
	cmd := &MessageSetType{}
	var keyLen int32
	err := binary.Read(r, binary.LittleEndian, &keyLen)
	if err != nil {
		return nil, err
	}
	cmd.Key = make([]byte, keyLen)
	err = binary.Read(r, binary.LittleEndian, &cmd.Key)
	if err != nil {
		return nil, err
	}
	var ValueLen int32
	err = binary.Read(r, binary.LittleEndian, &ValueLen)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.LittleEndian, &cmd.Value)
	if err != nil {
		return nil, err
	}
	return cmd, nil

}
