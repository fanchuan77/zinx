package znet

import (
	"zinx/src/zinx/ziface"
)

type Request struct {
	conn ziface.IConnection
	msg  ziface.IMessage
}

// GetConnection 得到当前连接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// GetMsgData 得到消息数据
func (r *Request) GetMsgData() []byte {
	return r.msg.GetMsgData()
}

// GetMsgId 得到消息
func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}

func NewRequest(conn ziface.IConnection, msg ziface.IMessage) ziface.IRequest {
	r := &Request{
		conn: conn,
		msg:  msg,
	}
	return r
}
