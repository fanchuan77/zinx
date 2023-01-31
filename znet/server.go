package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

var zinxLogo = `                                        
              ██                        
              ▀▀                        
 ████████   ████     ██▄████▄  ▀██  ██▀ 
     ▄█▀      ██     ██▀   ██    ████   
   ▄█▀        ██     ██    ██    ▄██▄   
 ▄██▄▄▄▄▄  ▄▄▄██▄▄▄  ██    ██   ▄█▀▀█▄  
 ▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀  ▀▀    ▀▀  ▀▀▀  ▀▀▀ 
                                        `
var topLine = `***********************************************`
var borderLine = `│`
var bottomLine = `***********************************************`

// Server IServer 的接口实现，定义一个Server的服务器模块
type Server struct {
	//服务器名称
	Name string
	//服务器监听IP版本
	IPVersion string
	//服务器监听IP
	IP string
	//服务器监听端口Port
	Port int
	//消息管理器
	MsgHandler ziface.IMsgHandle
	//连接管理器
	ConnMgr ziface.IConnManager
	//Server创建连接之后调用Hook函数
	OnConnStart func(conn ziface.IConnection)
	//Server销毁连接之前调用Hook函数
	OnConnStop func(conn ziface.IConnection)
}

// Start 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name:%s,Listenner at IP:%s:%d is starting \n",
		s.Name,
		s.IP,
		s.Port)
	fmt.Printf("[Zinx] Server Version:%s MaxConn:%d MaxPackageSize:%d \n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPackageSize)

	go func() {
		//获取一个TCP的Address
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve TCP addr err:", err)
			return
		}

		//监听服务器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "	err:", err)
			return
		}

		fmt.Println("start", s.Name, "succ,Listenning...")
		var cid uint32 = 0

		//阻塞的等待客户端连接，处理客户端请求
		for {
			//如果有客户端请求连接，阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			if s.ConnMgr.Len() >= uint32(utils.GlobalObject.MaxConn) {
				//TODO: 给客户端响应一个超出最大连接的错误包
				dp := NewDataPack()
				//将数据封装到消息包中并返回二进制数据
				binaryMsg, err := dp.Pack(NewMsgPackage(404, []byte("Too many connections!")))
				if err != nil {
					fmt.Println("Pack err", err)
					return
				}

				//向客户端发送二进制数据
				if _, err := conn.Write(binaryMsg); err != nil {
					fmt.Println("server send data err", err)
					return
				}

				err = conn.Close()
				if err != nil {
					fmt.Println("conn close err", err)
					return
				}
				continue
			}

			//将server注册的Router封装到新的连接对象
			//得到封装以后的Conn连接对象
			pakConn := NewConnection(s, conn, cid, s.MsgHandler)
			cid++

			//启动当前连接的业务处理
			go func() {
				err := pakConn.Start()
				if err != nil {
					fmt.Println("pakConn start err", err)
					return
				}
			}()
		}
	}()
}

// Serve 运行服务器
func (s *Server) Serve() {
	//启动Serve的服务功能
	s.Start()

	//TODO:启动服务后的额外业务拓展

	//阻塞状态
	select {}
}

// Stop 关闭服务器
func (s *Server) Stop() {
	//TODO:将一些连接的资源或信息回收
	fmt.Println("[STOP] zinx server name:", s.Name)

	//清空连接管理器
	s.ConnMgr.Clear()
}

// AddRouter 路由功能，给当前Server添加一个Router
func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
	fmt.Println("Add Router succ!!")
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

// SetOnConnStart 注册OnConnStart 钩子函数的方法
func (s *Server) SetOnConnStart(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

// SetOnConnStop 注册OnConnStop 钩子函数的方法
func (s *Server) SetOnConnStop(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

// CallOnConnStart 调用OnConnStart 钩子函数的方法
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("————> Call OnConnStart()...")
		s.OnConnStart(conn)
	}
}

// CallOnConnStop 调用OnConnStop 钩子函数的方法
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("————> Call OnConnStop()...")
		s.OnConnStop(conn)
	}
}

/*
初始化Server模块
*/
func NewServer() ziface.IServer {
	PrintLogo()
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}
	//启动工作池
	s.MsgHandler.StartWokerPool()

	return s
}

// PrintLogo 打印 Logo
func PrintLogo() {

	fmt.Println(zinxLogo)
	fmt.Println(topLine)
	fmt.Println(fmt.Sprintf("%s     [Github] https://github.com/aceld       %s", borderLine, borderLine))
	fmt.Println(bottomLine)
}
