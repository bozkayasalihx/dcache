package main

import (
	"bytes"
	"encoding/binary"
	"net"

	proto "github.com/bozkayasalih01x/cache/protocols"
)

func main() {

	con, err := net.Dial("tcp", ":3001")
	if err != nil {
		panic(err)
	}
	cmd := &proto.MessageSetType{
		Key:   []byte("foo"),
		Value: []byte("bar"),
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, proto.MSGSET)

	keyLen := int32(len(cmd.Key))
	binary.Write(buf, binary.LittleEndian, keyLen)
	binary.Write(buf, binary.LittleEndian, cmd.Key)

	valueLen := int32(len(cmd.Value))
	binary.Write(buf, binary.LittleEndian, valueLen)
	binary.Write(buf, binary.LittleEndian, cmd.Value)

	_, err = con.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}

	defer con.Close()

	select {}

}
