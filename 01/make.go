package main

import (
	"fmt"
	"strconv"
)

func test() {
	vals := make([]string, 10)

	for i := 0; i < 5; i++ {
		vals[i] = strconv.Itoa(i)
	}

	fmt.Println(vals, len(vals), cap(vals))
}

func main() {
	fmt.Println("========== make demo ==========")

	test()
}
