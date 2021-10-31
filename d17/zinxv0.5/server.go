package main

import (
	"dailytest/daily_test/d17/zinx/ziface"
	"dailytest/daily_test/d17/zinx/znet"
	"fmt"
)

//v0.5 解决TCP粘包问题，TLV type len value 封包格式来解决TCP粘包问题。
//第一次read得到datalen，根据datalen偏移读取消息数据
//第一次read得到datalen，根据datalen偏移读取消息数据
//zinx 解决消息的TLV序列化问题
//server 和 client stream 的形式传输数据的.
//包的形式12345,abcde这样一个包.
//对方需要知道数据有多长，避免处理粘包问题。
//消息长度，消息类型，消息内容。
//前两个固定八字节，后面的不确定
//5，200，“12345”
//6,201，abcdef，自定义的应用层协议
//
//
//

//基于zinx框架开发的服务器端程序
//V0.4 全局配置
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
	s:=znet.NewServer("[zinx v 0.5]")
	//2 给当前zinx 框架添加自定义router
	s.AddRouter(&PingRouter{})
	//3启动server
    s.Serve()
	//3

}
