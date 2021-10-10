package main

import (
	"fmt"
	"net"
)

type Server struct{
	Ip string
	Port int
}
//创建一个Server的接口
func NewServer(ip string,port int) *Server{
	server:=&Server{
		Ip:ip,
		Port:port,
	}
	return server
}
//业务处理模块
func (s *Server) Handler(conn net.Conn){
	//...当前链接的业务
	fmt.Println("链接建立成功")
}
//启动服务器的接口
func (s *Server)  Start(){
	//Socket listen
	listener,err:=net.Listen("tcp",fmt.Sprintf("" +
		"%s:%d",s.Ip,s.Port))
	if err!=nil{
		fmt.Println("net.Listen err:",err)
		return
	}
	//close listen socket
	defer listener.Close()
	for {
		//accept
		conn,err:=listener.Accept()
		//do handler
		if err!=nil{
			fmt.Println("listener accept err:",err)
			continue
		}
		go s.Handler(conn)
	}
	//close listen socket
}