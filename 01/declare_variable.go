package main

import "fmt"

func main() {
	fmt.Println("declare variable demo")
	test()
}

func test() {
	a := 10
	var b = 10
	var c int = 10
	fmt.Println(a, b, c)
}
