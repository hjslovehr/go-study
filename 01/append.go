package main

import (
	"fmt"
)

func test() {
	temp1 := []string{"1", "2", "3"}

	fmt.Println("temp1", temp1)

	temp2 := append(temp1, "4", "5")

	fmt.Println("temp2", temp2)
}

func main() {
	fmt.Println("========== Append Demo ==========")

	test()
}
