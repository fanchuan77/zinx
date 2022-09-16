package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lib/zinx/ziface"
)

/*
	存储所有Zinx框架需要的全局变量，供其他模块使用
	部分参数可以通过zinx.json油用户自行配置
*/
type GlobalObj struct {
	/*
		Server
	*/
	TcpServer        ziface.IServer //当前Zinx全局的Server对象
	Host             string         //当前服务器监听的IP
	TcpPort          int            //当前服务器监听的端口号
	Name             string         //当前服务器名称
	WorkerPoolSize   uint32         //当前业务工作池的 worker总量
	MaxWorkerTaskLen uint32         //Zinx框架允许单个消息队列包含任务的最大值
	/*
		Zinx
	*/
	Version        string //当前Zinx的版本号
	MaxConn        int    //当前服务器允许的最大连接数量
	MaxPackageSize uint32 //当前Zinx框架数据包的最大值
}

/*
	定义一个全局对外的GlobalObj
*/
var GlobalObject *GlobalObj

/*
	从zinx.json加载用户自定义的参数
*/
func (g *GlobalObj) Reload() {
	fmt.Println("Reload...")
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	//将json文件数据解析到GlobalObject中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

/*
	初始化GlobalObject对象
*/
func init() {
	fmt.Println("init...")
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerApp",
		Version:          "v0.8",
		Host:             "0.0.0.0",
		TcpPort:          8080,
		MaxConn:          1000,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
	}

	//从zinx.json加载用户自定义配置
	GlobalObject.Reload()
}
