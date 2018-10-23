package ipc

import (
	"testing"
	"fmt"
)

type EchoServer struct{}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{
		Code: "200",
		Body: "ECHO:" + method + params,
	}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1,_ := client1.Call("method1","1")
	resp2,_ := client1.Call("method2","2")
	fmt.Println(resp1)
	fmt.Println(resp2)
	if resp1.Body != "ECHO:method11" || resp2.Body != "ECHO:method22" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2:", resp2)
	}
	client1.Close()
	client2.Close()
}
