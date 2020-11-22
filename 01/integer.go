package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("integer demo")
	test()
	test2()
	test3()
}

func test() {
	var a int = 10
	fmt.Println(a)

	year := 2020
	fmt.Printf("year: %v\n", year)
	fmt.Printf("type year: %T\n", year)

	b := 1.0
	fmt.Printf("b: %v\n", b)
	fmt.Printf("type b: %T\n", b)

	var b1 float64 = 1.0
	fmt.Printf("b1: %v", b1)
	fmt.Printf("type b1: %T", b1)

	var c float32 = 1.0
	fmt.Printf("c: %v\n", c)
	fmt.Printf("type c: %T\n", c)
}

func test2() {
	a := 0x10
	fmt.Printf("0x%x\n", a)
	fmt.Printf("%v\n", a)
}

func test3() {
	var a uint8 = 1
	for {
		fmt.Println(a)
		a++
		time.Sleep(100 * time.Millisecond)
	}
}
