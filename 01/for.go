package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("for demo")
	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	test()
}

func test() {
	for {
		fmt.Printf("Now time is: %v\n", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}
