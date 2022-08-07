package main

import "zinx/src/zinx/znet"

func main() {
	s := znet.NewServer("[zinx_server]")
	s.Serve()
}
