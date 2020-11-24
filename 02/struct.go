package main

import (
	"fmt"
)

func test1() {
	var point struct {
		x float64
		y float64
	}

	point.x = 12.5
	point.y = 0.45

	fmt.Printf("%v, %v\n", point.x, point.y)
	fmt.Println(point)
}

// Student 自定义结构体
type Student struct {
	name string
	age  int
}

func test2() {
	var stu Student
	stu.name = "jim"
	stu.age = 20

	fmt.Println(stu)
	fmt.Println()

	stu2 := stu
	fmt.Println(stu2)
	fmt.Println()

	if stu == stu2 {
		fmt.Println("stu == stu2")
	} else {
		fmt.Println("stu != stu2")
	}
	fmt.Println()

	stu2.name = "tom"
	fmt.Println("stu", stu)
	fmt.Println("stu2", stu2)
	fmt.Println()

	if stu == stu2 {
		fmt.Println("stu == stu2")
	} else {
		fmt.Println("stu != stu2")
	}
	fmt.Println()

	stu3 := Student{name: "Yoyo", age: 20}
	fmt.Println(stu3)

	stu4 := Student{name: "KaKa"}
	fmt.Println(stu4)
}

func main() {
	fmt.Println("========== struct demo ==========")

	test1()

	fmt.Println()

	test2()
}
