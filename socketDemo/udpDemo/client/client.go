package main

import (
	"fmt"
	"net"
)

//udp客户端
func main() {
	//建立链接
	socket, err := net.DialUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 33333,
	}, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("链接服务器失败，err:", err)
		return
	}
	defer func() { _ = socket.Close() }()
	sendData := []byte("hello server")
	_, err = socket.Write(sendData) //发送数据
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}
	data := make([]byte, 4096)
	n, addr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败，err:", err)
		return
	}
	fmt.Printf("recv: %v, addr:%v, count:%v\n", string(data), addr, n)
}
