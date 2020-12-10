package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	test1()
}

func test1() {
	fmt.Println("Hello go")

	type Person struct {
		Name string
		Age  int
	}

	ss := `{
"Name": "king",
"Age": 20
}`
	data := []byte(ss)

	p := &Person{}

	err := json.Unmarshal(data, p)
	if err != nil {
		panic(err)
	}

	fmt.Println(*p)
}
