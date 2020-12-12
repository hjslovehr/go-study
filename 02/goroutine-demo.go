package main

import (
	"fmt"
	"time"
)

func test1() {
	fmt.Println("test begin")
	fmt.Println("test work...")
	time.Sleep(2 * time.Second)
	fmt.Println("test end")
}

func doSomthing(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " ...")
}

func test2() {
	for i := 0; i < 5; i++ {
		go doSomthing(i)
	}

	time.Sleep(4 * time.Second)
}

func main() {
	fmt.Println("---------- goroutine demo ----------")

	// go test1()
	// fmt.Println("main end")
	// time.Sleep(3 * time.Second)

	test2()
	fmt.Println("--- main end ---")
}
