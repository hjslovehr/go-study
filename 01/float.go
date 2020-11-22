package main

import (
	"fmt"
)

func main() {
	fmt.Println("float demo")
	test1()
	test2()
}

func test1() {
	a := 0.1
	b := 0.2
	fmt.Println(a + b)
	a = 1 / 3.0
	fmt.Println(a + a + a)
}

func test2() {
	a := 0.1
	a += 0.2
	fmt.Println(a == 0.3)
}
