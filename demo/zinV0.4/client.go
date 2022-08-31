//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"net"
	"time"
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

	//获取Server回显数据
	buf := make([]byte, 512)
	for {
		//向服务端发送数据
		_, err = conn.Write([]byte("Hello Zinx this is a Request!"))
		if err != nil {
			fmt.Println("client write data err", err)
			return
		}
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("client read data err", err)
			return
		}
		fmt.Printf("len:%d \n%s", cnt, buf)
		fmt.Println("--------------------")
		time.Sleep(1 * time.Second)
	}

}
