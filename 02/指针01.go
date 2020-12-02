package main

import "fmt"

func main() {
	fmt.Println("---------- 指针 01 ----------")
	test1()
	fmt.Println()
	test2()
	fmt.Println()
	test3()
}

func test1() {
	a := 10
	b := &a
	fmt.Println(b)
	fmt.Println(*b)

	var pstr *string
	str := "hello"
	pstr = &str

	fmt.Println(str)
	fmt.Println(pstr)
	fmt.Println(*pstr)

	*pstr = "world"
	fmt.Println(str)
	fmt.Println(pstr)
	fmt.Println(*pstr)
}

func test2() {
	type Person struct {
		name string
		age  int
	}

	tom := &Person{
		name: "Tom",
		age:  20,
	}

	fmt.Printf("tom: %v, %v", tom.name, tom.age)
}

func test3() {
	strs := &[3]string{"Hello", "World", "Jim"}

	fmt.Println(strs[0])
	fmt.Println(strs[1:2])

	ints := &[...]int{1, 2, 3, 4, 5, 6, 32, 2, 34, 5, 6, 7, 88, 68}

	fmt.Println(*ints)
	fmt.Println(ints[2:5])

	fmt.Println()
	bytes := []byte("abc")
	fmt.Println(bytes)
	fmt.Println(string(bytes))
}
