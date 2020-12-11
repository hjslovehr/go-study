package main

import (
	"bufio"
	"fmt"
	"net"
)

func start() {
	addr := "127.0.0.1:8888" // 表示本地所有ip的 8888 端口，也可以这样写：addr := ":8888"
	// 建立 tcp 服务
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("listen error:", err)
		panic(err)
	}
	defer listener.Close()

	fmt.Println("listen", addr, "start")

	for {
		// 等待客户端建立连接
		con, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			panic(err)
		}
		fmt.Println("Client", con.RemoteAddr())
		// 启动一个单独的 goroutine 去处理连接
		go process(con)
	}
}

func process(con net.Conn) {
	// 处理完关闭连接
	defer con.Close()

	// 针对当前连接做发送和接受操作
	for {
		reader := bufio.NewReader(con)
		var buf [1024]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("receive error:", err)
			return
		}

		data := string(buf[:n])
		fmt.Println("receive:", data)

		// 接受到数据后返回给客户端 OK
		_, err = con.Write([]byte("OK"))
		if err != nil {
			fmt.Println("send error:", err)
			return
		}
	}
}

func main() {
	fmt.Println("---------- socket tcp server demo ----------")
	start()
}
