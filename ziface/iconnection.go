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
	Send() error
}

//连接绑定的处理业务函数模型
type HandleFunc func(*net.TCPConn, []byte, int) error
