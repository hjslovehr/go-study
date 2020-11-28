package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Person struct
// 如果返回的结构字段名需要转换成其他格式，
// 可以用 `json:"your property name"` 来说明
// 如果是 xml，那么就是 `xml:"your property name"`
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func test() {
	p1 := Person{Name: "Tom", Age: 20}
	bytes, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		jsonStr := string(bytes)
		fmt.Println(jsonStr)
	}
}

func main() {
	fmt.Println("==================== struct demo2 ====================")

	test()
}
