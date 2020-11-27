package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Person struct
// 要用 json.Marshal 转 json，必须将属性名首字符大写
type Person struct {
	Name string
	Age  int
}

func test1() {
	xiaoming := Person{Name: "小明", Age: 15}
	fmt.Println(xiaoming)
	fmt.Printf("%+v\n", xiaoming)

	xiaozhang := Person{"小张", 20}
	fmt.Println(xiaozhang)
	fmt.Printf("%+v\n", xiaozhang)
}

func test2() {
	jim := Person{Name: "Jim", Age: 20}
	bytes, err := json.Marshal(jim)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(string(bytes))
	}
}

func main() {
	fmt.Println("==================== struct demo1 ====================")
	test1()
	test2()
}
