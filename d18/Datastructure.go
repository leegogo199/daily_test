package main

import (
	"fmt"
	"time"
)

type LinkNode struct{
	Data int64
	NextNode *LinkNode
}
func main(){
	var exitchan =make( chan bool,1)
	fmt.Println(exitchan)
	node:=new(LinkNode)
	node.Data=1
	node1:=new(LinkNode)
	node1.Data=2
	node.NextNode=node1
	node2:=new(LinkNode)
	node2.Data=3
	node1.NextNode=node2
	nowNode:=node
	for nowNode!=nil{
		fmt.Println(nowNode.Data)
		nowNode=nowNode.NextNode
	}
	a := make(chan bool,1)
	a <- true
	select {
	case <-a:
		{
			fmt.Println("ok")
		}
	case <-time.After(time.Second * 10):
		fmt.Println("10")
	default:
		fmt.Println("default")
	}

}