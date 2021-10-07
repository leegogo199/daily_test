package main

import (
	"flag"
	"fmt"
	"time"
	"golang.org/x/net/websocket"
)
var addr = flag.String("addr", "127.0.0.1:12345", "http service address")
func main() {
	flag.Parse()
	url := "ws://"+ *addr + "/ws"
	origin := "test://1111111/"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		fmt.Println(err)
	}
	go timeWriter(ws)

	for {
		var msg [512]byte
		_, err := ws.Read(msg[:])//此处阻塞，等待有数据可读
		if err != nil {
			fmt.Println("read:", err)
			return
		}

		fmt.Printf("received: %s\n", msg)
	}
}

func timeWriter(conn *websocket.Conn) {
	for {
		time.Sleep(time.Second * 2)
		websocket.Message.Send(conn, "hello world")
	}
}