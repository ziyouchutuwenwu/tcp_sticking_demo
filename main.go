package main

import (
	"fmt"
	"net"
	"tcp_sticking_demo/demo/go_tcp_client/header"
	"tcp_sticking_demo/demo/go_tcp_client/recv"
	"tcp_sticking_demo/demo/go_tcp_client/send"
)

func SendDemo() {
	data:= []byte("我没有来")
	pkgHeaderOption := header.GetPkgOptionWithHeaderSize(4)
	send.MakeDataToSend(data, pkgHeaderOption)
}

func ReadDemo(){
	conn, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Printf("dial failed, err:%v\n", err)
		return
	}
	recv.LoopRead(conn)
}

func main() {
	ReadDemo()
}