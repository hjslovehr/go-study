package main

import "fmt"

type doWork func()

func funcInvoke(work doWork) {
	work()
}

func test() {
	fmt.Println("test function invoke.")
}

func main() {
	fmt.Println("========== declare func type ==========")

	funcInvoke(test)

	f := func(s string) {
		fmt.Println(s)
	}

	f("hello function")

	func() {
		fmt.Println("function anonymous")
	}()
}
