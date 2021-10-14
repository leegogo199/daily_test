package main

import (
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
	}
	//链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client.conn = conn
	//返回对象
	return client
}

// 处理server 回应的消息，直接显示到标注输出即可
func (c *Client) DealResponse() {
	//一旦client.conn 有数据，就直接copy到stdout标注输出上，永久阻塞监听。
	io.Copy(os.Stdout, client.conn)
}

var serverIp string
var serverPort int

//./client -ip 127.0.0.1 -port 8888

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址（默认是127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口（默认是8888）")
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>链接服务器失败。。。")
		return
	}
	//单独开启一个goroutine 去处理server的回执消息
	go client.DealResponse()
	fmt.Println(">>>>链接服务器成功....")
	//启动客户端的业务
	select {}
}
