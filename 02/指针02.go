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

func (p *Person) birthday() {
	p.age++
}

func test2() {
	terry := &Person{
		name: "Terry",
		age:  15,
	}

	terry.birthday()
	fmt.Printf("%+v\n", terry)
}

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
}

func test3() {
	var board [8][8]rune
	reset(&board)

	fmt.Printf("%c", board[0][0])
}
