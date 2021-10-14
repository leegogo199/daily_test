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

func (c *Client) menu() bool {
	var key int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")
	fmt.Scanln(&key)

	if key >= 0 && key <= 3 {
		c.key = key
		return true
	} else {
		fmt.Println(">>>请输入合法范围内的数字<<<")
		return false
	}

}
func (c *Client)PublicChat(){
	//提示用户输入消息
	var chatMsg string
	fmt.Println(">>>>请输入聊天内容，exit退出。")
	fmt.Scanln("&chatMsg")
	//发给服务器
	for chatMsg!="exit"{
		//消息不为空则发送
		if len(chatMsg)!=0{
			sendMsg:=chatMsg+"\n"
			_,err:=c.conn.write([]byte(sendMsg))
			if err!=nil{
				fmt.Println("conn Write err:",err)
				break
			}
		}
		chatMsg=""
		fmt.Println(">>>>请输入聊天内容，exit退出。")
		fmt.Scanln(&chatMsg)
	}
}
func (c *Client) UpdateName() bool {
	fmt.Println(">>>>请输入用户名：")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + c.Name + "\n"
	_, err := c.conn.Write([byte](sendMsg))
	if err!=nil{
		fmt.Println("conn.Write err":err)
		return false
		
	}
	return true

}
func (c *Client) Run() {
	for c.key != 0 {
		for c.menu() != true {

		}
		//根据不同的模式处理不同的业务
		switch c.key {
		case 1:
			//公聊模式
			fmt.Println("公聊模式选择。。。")
			c.PublicChat()
			break
		case 2:
			//私聊模式
			fmt.Println("私聊模式选择。。。")
			break
		case 3:
			// 更新用户名
			fmt.Println("更新用户选择。。。")
			c.UpdateName()
			break
		}
	}

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
