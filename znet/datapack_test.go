package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

//测试datapack的封包、拆包功能
func TestDataPack(t *testing.T) {
	/*
		模拟服务器
	*/
	//创建 SocketTCP
	listenner, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	go func() {
		//从客户端读取数据，拆包处理
		for {
			//监听,获取连接
			conn, err := listenner.Accept()
			if err != nil {
				fmt.Println("accept err", err)
				return
			}
			go func(conn net.Conn) {
				//处理客户端的请求
				dp := NewDataPack()
				for {
					//进行第一次读取
					//先读取 Head (8字节)
					headData := make([]byte, dp.GetHeadLen())
					//读满 8字节的 Head信息
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("server read Head err", err)
						return
					}
					//拆包获得 MsgLen 和 ID
					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack Head err", err)
						return
					}

					//进行第二次读取
					if msgHead.GetMsgLen() > 0 {
						//从conn读，根据 head中的 MsgLen在读取 Data

						/*	使用Set方法也可以
							msg.SetMsgId(msgHead.GetMsgId())
							msg.SetMsgLen(msgHead.GetMsgLen())
						*/
						msg := msgHead.(*Message)

						//为 Message的 Data开辟足够大小的空间
						msg.Data = make([]byte, msg.GetMsgLen())

						//根据 MsgLen从io流中读取 Data
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("server read Data err", err)
							return
						}

						//读取完毕
						fmt.Println("MsgId:", msg.ID, "MsgLen:", msg.MsgLen, ",Data:", string(msg.Data))
					}
				}
			}(conn)
		}
	}()

	/*
		模拟客户端
	*/
	//客户端发起连接
	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println("client connect err", err)
		return
	}
	dp := NewDataPack()
	//封装第一个Msg
	msg1 := &Message{
		ID:     1,
		MsgLen: 4,
		Data:   []byte{'1', '2', '3', '4'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client connect err", err)
		return
	}
	//封装第二个Msg
	msg2 := &Message{
		ID:     2,
		MsgLen: 4,
		Data:   []byte{'4', '3', '2', '1'},
	}
	fmt.Println("pack...")
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client connect err", err)
		return
	}
	//将两个包粘在一起
	sendData1 = append(sendData1, sendData2...)
	//发送给服务器
	fmt.Println("send...")
	conn.Write(sendData1)
	//客户端阻塞
	select {}
}
