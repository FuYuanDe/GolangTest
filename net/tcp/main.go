// main.go
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var buf []byte

	SrvAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:3600")
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	} else {
		fmt.Println("resolve tcp addr success")
	}

	//创建连接
	conn, err := net.DialTCP("tcp", nil, SrvAddr)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	} else {
		fmt.Println("diag tcp success")
	}
	fmt.Println("sending msg to server")
	writeLen, err := conn.Write([]byte("hello server"))
	if err != nil {
		fmt.Errorf("%s", err.Error())
	} else {
		fmt.Println("-->sending msg to server success, msg length :%d", writeLen)
	}

	fmt.Println("enter reading...")
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}
	fmt.Println("server close")
	conn.Close()
	time.Sleep(2 * time.Second)
	return
}
