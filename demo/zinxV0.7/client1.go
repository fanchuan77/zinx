//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/src/zinx/znet"
)

func main() {
	fmt.Println("client start...")

	time.Sleep(1 * time.Second)

	//向Server请求连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("client connect err", err)
		return
	}

	for {
		dp := znet.NewDataPack()
		//将数据封装到消息包中并返回二进制数据
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(1, []byte("[zinx v0.7] client1 Test message:hello server...")))
		if err != nil {
			fmt.Println("Pack err", err)
			return
		}

		//向服务器发送二进制数据
		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("client send data err", err)
			return
		}

		//休息一秒,开始读取回显数据
		time.Sleep(1 * time.Second)

		//进行第一次读取
		//先读取 Head (8字节)
		headData := make([]byte, dp.GetHeadLen())
		//读满 8字节的 Head信息

		if _, err := io.ReadFull(conn, headData); err != nil {
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
			msg := msgHead.(*znet.Message)

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
		fmt.Println("--------------------")
		time.Sleep(1 * time.Second)
	}

}
