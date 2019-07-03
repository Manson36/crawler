package main

import (
	rpc2 "github.com/crawler/lang/maze/practicectice/jsonrpc/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpc2.DemoService{})

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
		}
		continue

		go jsonrpc.ServeConn(conn)
	}
}
