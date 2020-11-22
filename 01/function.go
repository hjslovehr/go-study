package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("========== funtion demo ==========")

	invokeTest(test)

	doWork(work, before, end)
}

func invokeTest(doSomething func()) {
	doSomething()
}

func test() {
	fmt.Println("test function invoke.")
}

func doWork(work func(int) string, before func(), end func()) {
	before()
	fmt.Println(work(5))
	end()
}

func work(count int) string {
	for i := 0; i < count; i++ {
		fmt.Println("Currnt time: ", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}

	return "word end"
}

func before() {
	fmt.Printf("before function invoke.\n")
}

func end() {
	fmt.Printf("end function invoke.\n")
}
