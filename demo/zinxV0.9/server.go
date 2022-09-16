package main

import (
	"fmt"
	"lib/zinx/ziface"
	"lib/zinx/znet"
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

//Test Hook
func Start(conn ziface.IConnection) {
	fmt.Println("--------------this is start()--------------")
}

func Stop(conn ziface.IConnection) {
	fmt.Println("--------------this is stop()--------------")
}

func main() {
	s := znet.NewServer()

	//注册Hook函数
	s.SetOnConnStart(Start)
	s.SetOnConnStop(Stop)

	//注册Router函数
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	s.Serve()
}
