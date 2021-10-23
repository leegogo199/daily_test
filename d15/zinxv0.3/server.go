package main

import (
	"dailytest/daily_test/d15/zinx/ziface"
	"dailytest/daily_test/d15/zinx/znet"
	"fmt"
)

//基于zinx框架开发的服务器端程序

//ping test 自定义路由
//创建一个server句柄，使用zinx的api
//给当前zinx框架添加一个自定义的router
//启动server
//需要继承baserouter，beforehandle，handle，afterhandle。
//
type  PingRouter struct{
	znet.BaseRouter
}
//test before
func (p *PingRouter) BeforeHandle(request ziface.IRequest){
	fmt.Println("Call Router BeforeHandle...")
	_,err:=request.GetConnection().GetTCPConnection().Write([]byte("before ping..\n"))
	if err!=nil{
		fmt.Println("call back ping error",err)
	}
}
//test handler
func (p *PingRouter) Handle(request ziface.IRequest){
	fmt.Println("Call Router Handle...")
	_,err:=request.GetConnection().GetTCPConnection().Write([]byte("ping .. ping..\n"))
	if err!=nil{
		fmt.Println("call  ping error",err)
	}

}
//test after
func (p *PingRouter) AfterHandle(request ziface.IRequest){
	fmt.Println("Call router after Handle...")
	_,err:=request.GetConnection().GetTCPConnection().Write([]byte("pong .. pong..\n"))
	if err!=nil{
		fmt.Println("call after pong error",err)
	}
}

func main(){
	//1创建一个server句柄
	s:=znet.NewServer("[zinx v 0.2]")
	//2 给当前zinx 框架添加自定义router
	s.AddRouter(&PingRouter{})
	//3启动server
    s.Serve()
	//3

}
