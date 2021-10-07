package main

import(
	"golang.org/x/net/websocket"
	"fmt"
	"net/http"
	"flag"
)

type WSServer struct {
	ListenAddr string
}

func (wss *WSServer)handler(conn *websocket.Conn){
	fmt.Printf("a new ws conn: %s->%s\n", conn.RemoteAddr().String(), conn.LocalAddr().String())
	var err error
	for {
		var reply string
		err = websocket.Message.Receive(conn, &reply)
		if err != nil {
			fmt.Println("receive err:",err.Error())
			break
		}
		fmt.Println("Received from client: " + reply)
		if err = websocket.Message.Send(conn, reply); err != nil {
			fmt.Println("send err:", err.Error())
			break
		}
	}
}
func (wss *WSServer)start()(error){
	http.Handle("/ws", websocket.Handler(wss.handler))
	fmt.Println("begin to listen")
	err := http.ListenAndServe(wss.ListenAddr, nil)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
		return err
	}
	fmt.Println("start end")
	return nil
}

func main(){
	addr  := flag.String("a", "127.0.1.1:12345", "websocket server listen address")
	flag.Parse()
	wsServer := &WSServer{
		ListenAddr : *addr,
	}
	wsServer.start()
	fmt.Println("------end-------")
}