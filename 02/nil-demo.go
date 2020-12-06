package main

import "fmt"

func main() {
	fmt.Println("---------- nil demo ----------")
	fmt.Println()
	// test1()
	fmt.Println()
	// test2()
	fmt.Println()
	test3()
	fmt.Println()
}

func test1() {
	var tp *int
	fmt.Println(tp)
	fmt.Println(*tp)
}

type person struct {
	name string
	age  int
}

func (p *person) birthday() {
	p.age++
}

func test2() {
	var pp *person
	fmt.Println(pp)
	pp.birthday()
}

func test3() {
	var fn func(a, b int) int
	fmt.Println(fn == nil)
}
