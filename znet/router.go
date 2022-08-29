package znet

import "zinx/src/zinx/ziface"

//实现Router之前先继承BaseRouter，再根据需要对方法进行重写
type BaseRouter struct {
}

//在处理conn业务之前的钩子方法Hook
func PreHandle(request ziface.IRequest) {

}

//在处理conn业务的主方法Hook
func Handle(request ziface.IRequest) {

}

//在处理conn业务之后的钩子方法Hook
func PostHandle(request ziface.IRequest) {

}
