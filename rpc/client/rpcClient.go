package main

import (
	"net"
	"net/rpc/jsonrpc"
	"GoStudy/rpc"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Div", rpc.Args{10, 3}, &result)
	fmt.Println(result, err)
	err = client.Call("DemoService.Div", rpc.Args{10, 0}, &result)
	fmt.Println(result, err)
}
