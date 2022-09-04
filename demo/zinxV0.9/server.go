//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"zinx/src/zinx/ziface"
	"zinx/src/zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("recv from client: msgId =", request.GetMsgId(),
		",data =", string(request.GetMsgData()))
	err := request.GetConnection().SendMsg(0, []byte("ping succ..."))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
}

//Test Handle
func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("recv from client: msgId =", request.GetMsgId(),
		",data =", string(request.GetMsgData()))
	err := request.GetConnection().SendMsg(1, []byte("Hello Zinx Router succ..."))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer()
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})
	s.Serve()
}
