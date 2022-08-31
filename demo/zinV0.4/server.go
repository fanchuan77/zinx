//go:build ignore
// +build ignore

package main

import (
	"zinx/src/zinx/ziface"
	"zinx/src/zinx/znet"
)

type PingRouter struct {
	BaseRouter znet.BaseRouter
}

//在处理conn业务之前的钩子方法Hook
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
}

//在处理conn业务的主方法Hook
func (this *PingRouter) Handle(request ziface.IRequest) {
	request.GetConnection().GetTCPConnection().Write([]byte("ping succ...\n"))

}

//在处理conn业务之后的钩子方法Hook
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
}

func main() {
	s := znet.NewServer()
	s.AddRouter(&PingRouter{})
	s.Serve()
}
