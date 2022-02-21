package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//1.建立与服务端的链接
//2.进行数据收发
//3.关闭链接
func main() {
	//建立链接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Connect to server failed, err:", err)
		return
	}
	//延时调用，关闭链接
	defer func() { _ = conn.Close() }()

	//等待从标准输入读取数据, 读取到 *Reader 中
	inputReader := bufio.NewReader(os.Stdin)
	for {
		//循环读取用户输入，以换行为结束符
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { //如果输入q,就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo)) //发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
