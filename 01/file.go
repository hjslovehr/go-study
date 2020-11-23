package main

import (
	"fmt"
	"io"
	"os"
)

func fileCreate(fileName string) {
	// 创建文件，会覆盖已存在的文件，并且已存在的文件内容会被清空
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Create file error")
		panic(err)
	} else {
		fmt.Println("Create file " + fileName + " successed.")

		defer file.Close()
	}
}

func writeFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file fialed.")
		panic(err)
	} else {

		defer file.Close()

		// writer := bufio.NewWriter(file)
		// writer.WriteString("Hello, this line is program write.\n")

		// writer.Flush()

		file.WriteString("Hello, this line is program write.\n")

		fmt.Println("write file successed.")
	}
}

func fileRead(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file errir")
		panic(err)
	} else {
		defer file.Close()

		bytes := make([]byte, 1024)

		for {
			len, err := file.Read(bytes)
			if err != nil && err != io.EOF {
				fmt.Println(err)
			}

			if len == 0 {
				break
			}

			fmt.Println(string(bytes))
		}
	}
}

func main() {
	fmt.Println("==================== file demo ====================")

	// fileCreate("txtfile.txt")

	writeFile("txtfile.txt")

	fileRead("txtfile.txt")
}
