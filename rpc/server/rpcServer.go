package main

import (
	"net/rpc"
	rpc2 "GoStudy/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)

//通过telnet localhost 1234 连接，输入：{"method":"DemoService.Div","params":[{"A":3,"B":4}],"id":1}
//调用成功
func main() {
	rpc.Register(rpc2.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("acceppt error:%v", err)
		}
		jsonrpc.ServeConn(conn)
	}

}
