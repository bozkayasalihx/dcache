package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"

	proto "github.com/bozkayasalih01x/cache/protocols"
)

func main() {

	conn, err := net.Dial("tcp", ":3001")
	if err != nil {
		panic(err)
	}
	// setCmd := &proto.MessageSetType{
	// 	Key:   []byte("foo"),
	// 	Value: []byte("bar"),
	// }

	// setBuff := new(bytes.Buffer)
	// binary.Write(setBuff, binary.LittleEndian, proto.MSGSET)

	// keyLen := int32(len(setCmd.Key))
	// binary.Write(setBuff, binary.LittleEndian, keyLen)
	// binary.Write(setBuff, binary.LittleEndian, setCmd.Key)

	// valueLen := int32(len(setCmd.Value))
	// binary.Write(setBuff, binary.LittleEndian, valueLen)
	// binary.Write(setBuff, binary.LittleEndian, setCmd.Value)

	// _, err = conn.Write(setBuff.Bytes())
	// if err != nil {
	// 	panic(err)
	// }

	// time.Sleep(time.Second * 5)
	// getCmd := &proto.MessageGetType{
	// 	Key: []byte("foo"),
	// }

	// getBuff := new(bytes.Buffer)
	// binary.Write(getBuff, binary.LittleEndian, proto.MSGGET)
	// getkeylen := int32(len(getCmd.Key))
	// binary.Write(getBuff, binary.LittleEndian, getkeylen)
	// binary.Write(getBuff, binary.LittleEndian, getCmd.Key)

	// _, err = conn.Write(getBuff.Bytes())
	// if err != nil {
	// 	panic(err)
	// }

	// bufData := make([]byte, 1024)
	// n, err := conn.Read(bufData)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("make me alive -> %v\n", string(bufData[:n]))

	WriteSomeTestData(conn)
	go ReadSomeTestData(conn)
	defer conn.Close()

	select {}

}

var (
	iterate = 10
)

func WriteSomeTestData(conn net.Conn) {
	for i := 0; i < iterate; i++ {
		fmt.Println("writing new test data to connection")

		setCmd := &proto.MessageSetType{
			Key:   []byte(fmt.Sprintf("key_%d", i)),
			Value: []byte(fmt.Sprintf("val_%d", i)),
		}

		setBuff := new(bytes.Buffer)
		binary.Write(setBuff, binary.LittleEndian, proto.MSGSET)

		keyLen := int32(len(setCmd.Key))
		binary.Write(setBuff, binary.LittleEndian, keyLen)
		binary.Write(setBuff, binary.LittleEndian, setCmd.Key)

		valueLen := int32(len(setCmd.Value))
		binary.Write(setBuff, binary.LittleEndian, valueLen)
		binary.Write(setBuff, binary.LittleEndian, setCmd.Value)

		_, err := conn.Write(setBuff.Bytes())
		if err != nil {
			panic(err)
		}

		time.Sleep(time.Millisecond * 200)

	}

}

func ReadSomeTestData(conn net.Conn) {
	fmt.Println("reading new test data from connection")
	for i := 0; i < iterate; i++ {
		getCmd := &proto.MessageGetType{
			Key: []byte(fmt.Sprintf("key_%d", i)),
		}

		getBuff := new(bytes.Buffer)
		binary.Write(getBuff, binary.LittleEndian, proto.MSGGET)
		getkeylen := int32(len(getCmd.Key))
		binary.Write(getBuff, binary.LittleEndian, getkeylen)
		binary.Write(getBuff, binary.LittleEndian, getCmd.Key)

		_, err := conn.Write(getBuff.Bytes())
		if err != nil {
			panic(err)
		}

		bufData := make([]byte, 1024)
		n, err := conn.Read(bufData)
		if err != nil {
			panic(err)
		}

		fmt.Printf("make me alive -> %v\n", string(bufData[:n]))
	}

}
