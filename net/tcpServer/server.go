// server.go
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	LocalAddr, err := net.ResolveTCPAddr("tcp", ":3600")
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	} else {
		fmt.Println("resolve local addr success")
	}

	//conn, err := net.DialTCP("tcp",)
	listener, err := net.ListenTCP("tcp", LocalAddr)
	if err != nil {
		fmt.Println("%s", err.Error())
		return
	} else {
		fmt.Println("-->Listen tcp success")
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Errorf("%s", err.Error())
		} else {
			fmt.Println("-->accept one")
			go handleTcpClient(conn)
		}
	}
}

func handleTcpClient(conn *net.TCPConn) {
	var buf []byte = make([]byte, 100, 1000)

	conn.SetKeepAlivePeriod(time.Second * 1)
	conn.SetKeepAlive(true)
	fmt.Println("-->reading now")
	readLen, err := conn.Read(buf)
	if err != nil {
		fmt.Errorf("%s", err.Error())
	} else {
		fmt.Printf("receive msg length :%d, msg :%s", readLen, string(buf))
	}
	time.Sleep(2 * time.Second)
	conn.Close()
}
