package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() //处理完毕之后关闭连接
	//针对当前连接做数据处理
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed,err:%v\n", err)
			break
		}
		recv := string(buf[:n])
		fmt.Println("接收到的数据：", recv)
		conn.Write([]byte("ok"))
	}
}

func Tcpserverinit() {
	// 1. 开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		return
	}
	fmt.Printf("启动tcp服务端成功,ip:%s,port:%d\n", "127.0.0.1", 20000)
	for {
		// 等待客户端来建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed,err:%v\n", err)
			continue
		}
		// 3. 启动一个单独的goroutine去处理连接
		go process(conn)
	}
}
