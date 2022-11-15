package znet

import "zinx/ziface"

// BaseRouter 实现 Router之前先继承BaseRouter，再根据需要对方法进行重写
type BaseRouter struct {
}

// PreHandle 在处理conn业务之前的钩子方法Hook
func (bp *BaseRouter) PreHandle(request ziface.IRequest) {

}

// Handle 在处理conn业务的主方法Hook
func (bp *BaseRouter) Handle(request ziface.IRequest) {

}

// PostHandle 在处理conn业务之后的钩子方法Hook
func (bp *BaseRouter) PostHandle(request ziface.IRequest) {

}
