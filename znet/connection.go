package znet

import (
	"fmt"
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
	//当前连接绑定的业务处理函数
	handleAPI ziface.HandleFunc
	//告知当前连接已经退出的channel
	ExitChan chan bool
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

//发送数据
func (c *Connection) Send() error {
	return nil
}

//连接的读业务函数
func (c *Connection) startReader() error {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("ConnID =", c.ConnID, "Reader is exit,Remote Addr is", c.RemoteAddr())
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}
		fmt.Printf("Reader read:%s \n", buf)
		//调用连接所绑定的业务函数
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID", c.ConnID, "handle is err", err)
			break
		}
	}
	return nil
}

//启动连接
func (c *Connection) Start() error {
	fmt.Println("conn Start()... ConnID:", c.ConnID)
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

func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) ziface.IConnection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		handleAPI: callback_api,
		ExitChan:  make(chan bool, 1),
	}
	return c
}
