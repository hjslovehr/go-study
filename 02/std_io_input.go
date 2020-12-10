package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func test() {
	fmt.Println("You can input something and input \".exit\" key end.")

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if line == ".exit" {
			return
		}
		fmt.Println("You input:", line)
	}
}

func test2() {
	fmt.Println("You can input something and input \".exit\" key end.")

	input := bufio.NewReader(os.Stdin)
	for {
		s, err := input.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		if strings.TrimSpace(s) == ".exit" {
			return
		}
		fmt.Println("You input:", s)
	}
}

func main() {
	fmt.Println("---------- standard IO input demo ----------")
	// test()
	test2()
}
