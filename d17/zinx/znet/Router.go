package znet

import "dailytest/daily_test/d17/zinx/ziface"

// 实现router时，先嵌入这个BaseRouter基类，然后根据这个基类的方法进行重写。

type BaseRouter struct{

}
// baserouter的方法都为空，router默认 为空。
//处理前的方法
func (br *BaseRouter)BeforeHandle(request ziface.IRequest){}
//处理时的方法
func (br *BaseRouter)Handle(request ziface.IRequest){}
//处理后的方法
func (br *BaseRouter)AfterHandle(request ziface.IRequest){}
