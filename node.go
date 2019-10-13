package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"

	b "github.com/AdyRehman/A02"
)

func handleConnection(conn net.Conn) {
	chainhead := &b.Block{}
	log.Println("A client has connected ", conn.RemoteAddr())
	dec := gob.NewDecoder(conn)
	err := dec.Decode(chainhead)
	if err != nil {
		log.Println(err)
	}
	b.ListBlocks(chainhead)
}

var chainhead *b.Block

const Satoshi_Address = "localhost:6001"
const Satoshi_Address2 = "localhost:6002"

type Node_Info struct {
	Name string
	Port string
}

func main() {
	conn, err := net.Dial("tcp", Satoshi_Address)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	log.Println("client has connected ", conn.RemoteAddr())
	node_info := &Node_Info{os.Args[1], os.Args[2]}
	gobEncoder := gob.NewEncoder(conn)
	err = gobEncoder.Encode(node_info)
	if err != nil {
		log.Println(err)
	}
	ln, err := net.Listen("tcp", ":6002")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
