package main

import "dailytest/daily_test/d14/zinx/znet"

//基于zinx框架开发的服务器端程序


func main(){
	//1创建一个server句柄
	s:=znet.NewServer("[zinx v 0.2]")
	//2启动server
    s.Serve()
	//3



}
