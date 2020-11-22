package main

import (
	"fmt"
	"strings"
)

func test(params []string) {
	fmt.Println(strings.Join(params, ""))

}

func main() {
	fmt.Println("========== Slice Demo ==========")

	temp := []string{"hello", " ", "world"}

	test(temp)
}
