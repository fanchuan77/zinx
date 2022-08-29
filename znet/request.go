package znet

import (
	"zinx/src/zinx/ziface"
)

type Request struct {
	conn ziface.IConnection
	data []byte
}

//得到当前连接
func (r *Request) GetConnection() ziface.IConnection {

	return r.conn
}

//得到请求数据
func (r *Request) GetData() []byte {

	return r.data
}

func NewRequest(conn ziface.IConnection, data []byte) ziface.IRequest {
	r := &Request{
		conn: conn,
		data: data,
	}
	return r
}
