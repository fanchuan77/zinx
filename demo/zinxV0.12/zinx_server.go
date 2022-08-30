//go:build ignore
// +build ignore

package main

import "zinx/src/zinx/znet"

/*
	master branch
*/
func main() {
	s := znet.NewServer("[zinx_server]")
	s.Serve()
}
