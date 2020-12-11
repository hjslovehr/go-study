package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func start() {
	addr := "127.0.0.1:8888" // 表示本地所有ip的 8888 端口，也可以这样写：addr := ":8888"
	// 1、与服务端建立连接
	con, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("connect error:", err)
		panic(err)
	}

	defer con.Close()

	fmt.Println("Connect", addr)

	// 2、使用 con 连接进行数据的发送和接收
	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("You input: ")
		s, err := input.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}

		_, err = con.Write([]byte(s))
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		var buf [1024]byte
		n, err := con.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed:%v\n", err)
			return
		}
		fmt.Printf("receive: %v\n", string(buf[:n]))
	}
}

func main() {
	fmt.Println("---------- socket tcp client demo ----------")
	start()
}
