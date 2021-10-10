package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct{
	Ip string
	Port int
	//在线用户的列表
	OnlineMap map[string]*User
	maplock sync.RWMutex
	// 消息广播的channel
	Message chan string
}
//创建一个Server的接口
func NewServer(ip string,port int) *Server{
	server:=&Server{
		Ip:ip,
		Port:port,
		OnlineMap:make(map[string]*User),
		Message: make(chan string),
	}
	return server
}
//监听Message广播消息Channel的goroutine，
//一旦有消息就发送给全部的在线user。
func (s *Server) ListenMessager(){
	for {
		msg:=<-s.Message
		//将msg发送给全部的在线user
		s.maplock.Lock()
		for _,cli:=range s.OnlineMap{
			cli.C<-msg
		}
		s.maplock.Unlock()
	}
}
// 广播消息的方法
func (s *Server)BroadCast(user *User,msg string){
	sendMsg:="["+user.Addr+"]"+user.Name+":"+msg
	s.Message<-sendMsg

}

//业务处理模块
func (s *Server) Handler(conn net.Conn){
	//...当前链接的业务
	//fmt.Println("链接建立成功")
	//用户上线，将用户加入到onlinemap中。
     user:=NewUser(conn)
	// 广播当前用户上线消息
	s.maplock.Lock()
     s.OnlineMap[user.Name]=user
	s.maplock.Unlock()
     s.BroadCast(user,"已上线")
     // 当前handler 阻塞
     fmt.Println(user.Name,"已上线")
     select{

	 }

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
	//启动监听Message 的goroutine
	go s.ListenMessager()
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