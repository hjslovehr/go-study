package main

import "fmt"

func main() {
	fmt.Println("---------- 指针 02 ----------")
	test1()
	fmt.Println()
	test2()
	fmt.Println()
	test3()
}

// Person struct
type Person struct {
	name string
	age  int
}

func add(p *Person) {
	p.age++
}

func test1() {
	p := &Person{name: "Jim", age: 20}

	fmt.Println(*p)

	add(p)

	fmt.Println(*p)
}

func test2() {

}

func test3() {

}
