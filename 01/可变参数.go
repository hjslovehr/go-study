package main

import "fmt"

func test(prefex string, args ...string) []string {
	newArgs := args[:]

	for i := 0; i < len(newArgs); i++ {
		newArgs[i] = prefex + "-" + newArgs[i]
	}

	return newArgs
}

func main() {
	fmt.Println("========== 可变参数 ==========")

	ss := []string{"abc", "def", "ghi"}
	fmt.Println(ss)

	newss := test("new", ss...)
	fmt.Println(newss)
}
