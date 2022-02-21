package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer func() { _ = conn.Close() }() //关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) //读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client发来的数据：", recvStr)
		respStr := "hello, I'm server!!"
		_, _ = conn.Write([]byte(respStr)) //发送数据
	}

}

//1.监听端口
//2.接收客户端请求建立链接
//3.创建goroutine处理链接。
func main() {
	//1. 创建监听 -- 监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer func() { _ = listen.Close() }()
	for {
		//2. 等待客户端建立链接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		//3. 并发处理链接
		go process(conn)

	}

}
