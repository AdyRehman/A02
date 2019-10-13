package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"

	b "github.com/AdyRehman/A02"
)

const Self_Port = ":6001"
const Self_Port2 = ":6002"

type Node_Info struct {
	Name string
	Port string
}

var chainhead *b.Block

//Node_Info_list := []Node_Info{}
var names []string
var ports []string

func main() {
	min_nodes_x := 0
	//Node_Info_list := []Node_Info{}
	ln, err := net.Listen("tcp", Self_Port)
	if err != nil {
		log.Fatal(err)
	}

	// This loop should run for iterations = totalNodes
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// routine for handling initial node connections
		if min_nodes_x <= 2 {
			go handleInitialConnection(conn)
			min_nodes_x = min_nodes_x + 1
		} else {
			conn, err := net.Dial("tcp", Self_Port2)
			if err != nil {
				fmt.Println(err)
			} else {
				log.Println("client has connected ", conn.RemoteAddr())
				gobEncoder := gob.NewEncoder(conn)
				err = gobEncoder.Encode(chainhead)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func handleInitialConnection(conn net.Conn) {
	log.Println("A client has connected ", conn.RemoteAddr())
	var recvd_Node_Info Node_Info
	dec := gob.NewDecoder(conn)
	err := dec.Decode(&recvd_Node_Info)
	if err != nil {
		log.Println(err)
	}
	names = append(names, recvd_Node_Info.Name)
	ports = append(names, recvd_Node_Info.Port)
	chainhead = b.InsertBlock("paySatoshi100", chainhead)
	b.ListBlocks(chainhead)
	fmt.Println(recvd_Node_Info.Name, recvd_Node_Info.Port)
}

func handleConnection(conn net.Conn) {
	log.Println("A client has connected ", conn.RemoteAddr())
	gobEncoder := gob.NewEncoder(conn)
	err := gobEncoder.Encode(chainhead)
	if err != nil {
		log.Println(err)
	}
	b.ListBlocks(chainhead)
}
