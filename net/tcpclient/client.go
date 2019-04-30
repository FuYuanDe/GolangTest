// main.go
package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	//"syscall"
	"errors"
	"time"
)

func GetPort(addr string) (uint16, error) {
	if len(addr) == 0 {
		return 0, errors.New("invalid addr")
	} else {
		offset := strings.LastIndex(addr, ":")
		if offset == -1 {
			return 0, errors.New("invalid addr, no port found")
		} else {
			data := addr[(offset + 1):]
			port, err := strconv.ParseUint(string(data), 10, 16)
			if err != nil {
				return 0, err
			} else {
				return uint16(port), nil
			}
		}
	}
}

func GetIP(addr string) (string, error) {
	if len(addr) == 0 {
		return "", errors.New("invalid addr")
	} else {
		offset := strings.LastIndex(addr, ":")
		if offset == -1 {
			return "", errors.New("invalid addr, no port found")
		} else {
			data := addr[:offset]
			return string(data), nil
		}
	}
}

func main() {
	buffer := make(chan []byte, 100)

	SrvAddr, err := net.ResolveTCPAddr("tcp", "172.25.1.192:49")
	//SrvAddr, err := net.ResolveTCPAddr("tcp", "172.25.100.240:3600")
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
	//LocalAddr, err := syscall.Getsockname(conn.LocalAddr())
	addr := conn.LocalAddr()
	port, err := GetPort(addr.String())
	if err != nil {
		fmt.Printf("GetPort fail,%s\n", err.Error())
	} else {
		fmt.Println(port)
	}

	ip, err := GetIP(addr.String())
	if err != nil {
		fmt.Printf("GetIP fail,%s\n", err.Error())
	} else {
		fmt.Printf("IP:%s", ip)
	}
	go func() {
		//Buf := make([]byte, 1024)
		Buf := []byte("c00l010000000000000000000000000")
		for {
			select {
			case _ = <-buffer:
				//case buf := <-buffer:
				//copy(Buf, buf)
				//data := Buf[:len(buf)]

				//writeLen, err := conn.Write(buf[0:])
				writeLen, err := conn.Write(Buf)
				if err != nil {
					fmt.Errorf("%s", err.Error())
				} else {
					fmt.Printf("-->sending msg to server success, msg length :%d", writeLen)
				}
			}
		}
	}()
	var count int
	for {
		if count > 2 {
			break
		} else {
			count++
		}
		fmt.Println("send to server")
		buf := []byte("hello server")
		buffer <- buf
		time.Sleep(time.Second * 3)
	}

	conn.Close()
	time.Sleep(20 * time.Second)
	return
}
