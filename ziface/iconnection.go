package ziface

import (
	"net"
)

type IConnection interface {
	// Start 启动连接
	Start() error

	// Stop 停止连接
	Stop() error

	// GetTCPConnection 获取当前连接的conn对象
	GetTCPConnection() *net.TCPConn

	// GetConnID 获取连接ID
	GetConnID() uint32

	// RemoteAddr 获取客户端连接地址及端口
	RemoteAddr() net.Addr

	// SendMsg 发送数据
	SendMsg(msgId uint32, data []byte) error

	// SetProperty 设置连接属性
	SetProperty(key string, value interface{})

	// GetProperty 获取连接属性
	GetProperty(key string) (interface{}, error)

	// RemoveProperty 移除连接属性
	RemoveProperty(key string)
}
