package main

import "fmt"

func main() {
	fmt.Println("========== Array Demo ==========")

	var color [3]string
	color[0] = "red"
	color[1] = "green"
	color[2] = "yellow"

	fmt.Println(len(color))

	fmt.Println(color)

	ints := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(ints)

	for i := 0; i < len(ints); i++ {
		fmt.Println(i, ":", ints[i])
	}
	fmt.Println()

	for i, v := range ints {
		fmt.Println(i, ":", v)
	}
}
