package main

import (
	"fmt"
	"github.com/crawler/lang/rpc2"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	Client := jsonrpc.NewClient(conn)

	var result float64

	err = Client.Call("DemoService.Div", rpc2.Args{A:2, B:3}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
