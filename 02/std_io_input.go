package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("---------- standard IO input demo ----------")
	test()
}

func test() {
	fmt.Println("You can input something and input \".exit\" key end.")

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if line == ".exit" {
			break
		}
		fmt.Println("You input:", line)
	}
}
