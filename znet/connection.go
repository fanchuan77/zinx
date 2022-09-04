package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"zinx/src/zinx/ziface"
)

type Connection struct {
	//当前连接所属Server
	TcpServer ziface.IServer

	//当前连接的conn对象
	Conn *net.TCPConn

	//当前连接ID
	ConnID uint32

	//当前连接状态
	isClosed bool

	//告知当前连接已经退出的channel
	ExitChan chan bool

	//用于读、写 Goroutine的消息通信
	MsgChan chan []byte

	//当前连接处理的Router方法
	MsgHandler ziface.IMsgHandle

	//连接属性
	property map[string]interface{}

	//连接书的的保护锁
	propertyLock sync.RWMutex
}

//获取连接ID
func (c *Connection) GetConnID() uint32 {
	//fmt.Println("conn get connID", c.ConnID, "succ...")
	return c.ConnID
}

//获取当前连接的conn对象
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取客户端连接地址及端口
func (c *Connection) RemoteAddr() net.Addr {
	//fmt.Println("conn get RemoteAddr", c.Conn.RemoteAddr(), "succ...")
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
	//消息发送给 channel,Writer读取后写入连接
	c.MsgChan <- binaryMsg
	return nil
}

//读消息的 Goroutine
func (c *Connection) startReader() error {
	fmt.Println("[Reader Goroutine is running...]")
	defer fmt.Println(c.RemoteAddr().String(), "[Reader Goroutine exit!]")
	defer c.Stop()

	for {

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
		//将request传入消息队列
		go c.MsgHandler.SendMsgToTaskQueue(&req)
	}
}

//写消息的 Goroutine,负责发送消息给客户端
func (c *Connection) startWriter() {
	fmt.Println("[Writer Goroutine is running..]")
	defer fmt.Println(c.RemoteAddr().String(), "[Writer Goroutine exit!]")

	//阻塞等待消息 channel,进行写消息给客户端
	for {
		select {
		case data := <-c.MsgChan:
			//有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data err", err)
				return
			}
		case <-c.ExitChan:
			//代表 Reader已经退出,此时 Writer也要退出
			return
		}
	}
}

//启动连接
func (c *Connection) Start() error {
	fmt.Println("start connection succ.. ConnID:", c.ConnID)

	//启动连接的读业务函数
	go c.startReader()

	//启动连接的写业务函数
	go c.startWriter()

	//调用Hook函数
	c.TcpServer.CallOnConnStart(c)

	return nil
}

//停止连接
func (c *Connection) Stop() error {
	fmt.Println("ConnID:", c.ConnID, "Connection exit")
	if c.isClosed {
		return nil
	}
	//关闭连接
	err := c.Conn.Close()
	if err != nil {
		fmt.Println("ConnID:", c.ConnID, "Connection exit err", err)
	}
	//通知 channel Reader已退出
	c.ExitChan <- true
	c.isClosed = true

	//调用Hook函数
	c.TcpServer.CallOnConnStop(c)

	//回收资源
	close(c.MsgChan)
	close(c.ExitChan)

	//将当前Connection从连接管理器移除
	c.TcpServer.GetConnMgr().Remove(c)

	return nil
}

//设置连接属性
func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[key] = value
}

//获取连接属性
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if value, ok := c.property[key]; ok {
		return value, nil
	}
	return nil, errors.New("Connection Property Not FOUND")
}

//移除连接属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}

func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) ziface.IConnection {
	c := &Connection{
		TcpServer:  server,
		Conn:       conn,
		ConnID:     connID,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
		MsgChan:    make(chan []byte),
		MsgHandler: msgHandler,
		property:   make(map[string]interface{}),
	}

	//将Connection加入连接管理器
	c.TcpServer.GetConnMgr().Add(c)

	return c
}
