package ziface

import (
	"net"
)

type IConnection interface {
	//启动连接
	Start() error

	//停止连接
	Stop() error

	//获取当前连接的conn对象
	GetTCPConnection() *net.TCPConn

	//获取连接ID
	GetConnID() uint32

	//获取客户端连接地址及端口
	RemoteAddr() net.Addr

	//发送数据
	SendMsg(msgId uint32, data []byte) error

	//设置连接属性
	SetProperty(key string, value interface{})

	//获取连接属性
	GetProperty(key string) (interface{}, error)

	//移除连接属性
	RemoveProperty(key string)
}
