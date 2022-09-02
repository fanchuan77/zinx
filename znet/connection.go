package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/src/zinx/ziface"
)

type Connection struct {
	//当前连接的conn对象
	Conn *net.TCPConn
	//当前连接ID
	ConnID uint32
	//当前连接状态
	isClosed bool
	//告知当前连接已经退出的channel
	ExitChan chan bool
	//当前连接处理的Router方法
	MsgHandler ziface.IMsgHandle
}

//获取连接ID
func (c *Connection) GetConnID() uint32 {
	fmt.Println("conn get connID", c.ConnID, "succ...")
	return c.ConnID
}

//获取当前连接的conn对象
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取客户端连接地址及端口
func (c *Connection) RemoteAddr() net.Addr {
	fmt.Println("conn get RemoteAddr", c.Conn.RemoteAddr(), "succ...")
	return c.Conn.RemoteAddr()
}

//发送数据,将数据先封包,再发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	dp := NewDataPack()
	msgPack := NewMsgPackage(msgId, data)

	//获得要发送给客户端的二进制数据
	binaryMsg, err := dp.Pack(msgPack)
	if err != nil {
		fmt.Println("Pack error msg id =", msgId)
		return errors.New("Pack msg error")
	}
	_, err = c.Conn.Write(binaryMsg)
	if err != nil {
		fmt.Println("write error msg id =", msgId)
		return errors.New("write msg error")
	}
	return nil
}

//连接的读业务函数
func (c *Connection) startReader() error {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("ConnID =", c.ConnID, "Reader is exit,Remote Addr is", c.RemoteAddr())
	defer c.Stop()

	for {
		//读取客户端数据到 buf
		// buf := make([]byte, 512)
		// _, err := c.Conn.Read(buf)
		// if err != nil {
		// 	fmt.Println("recv buf err", err)
		// 	continue
		// }
		// fmt.Printf("Reader read:%s \n", buf)

		//获取消息封装模块对象
		dp := NewDataPack()

		//进行第一次读取
		//读取客户端 Msg Head 二进制流 (8个字节)
		headData := make([]byte, dp.GetHeadLen())
		//读满 8字节的 Head信息
		_, err := io.ReadFull(c.Conn, headData)
		if err != nil {
			fmt.Println("server read Msg Head err", err)
			return err
		}

		//拆包,得到 MsgLen 和 ID
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack Msg Head err", err)
			return err
		}

		//进行第二次读取
		//根据 MsgLen 再次读取 Data
		if msg.GetMsgLen() > 0 {
			//从conn读，根据 MsgLen读取 Data
			//为 msg的 Data开辟足够大小的空间
			Data := make([]byte, msg.GetMsgLen())

			//根据 MsgLen从io流中读取 Data
			_, err := io.ReadFull(c.Conn, Data)
			if err != nil {
				fmt.Println("server read Msg Data err", err)
				return err
			}

			//Data放入 message
			msg.SetMsgData(Data)
		}

		//封装一个Request对象
		req := Request{
			conn: c,
			msg:  msg,
		}
		//执行当前连接绑定的Router方法
		go c.MsgHandler.DoMsgHandler(&req)
	}
}

//启动连接
func (c *Connection) Start() error {
	fmt.Println("start connection succ.. ConnID:", c.ConnID)
	//启动连接的读业务函数
	go c.startReader()

	//TODO:启动从当前连接写数据的业务函数

	return nil
}

//停止连接
func (c *Connection) Stop() error {
	fmt.Println("conn Stop()...ConnID:", c.ConnID)
	if c.isClosed {
		return nil
	}
	c.isClosed = true
	return nil
}

func NewConnection(conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) ziface.IConnection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
		MsgHandler: msgHandler,
	}
	return c
}
