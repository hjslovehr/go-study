package main

import "fmt"

func main() {
	for i := 'A'; i < 'A'+26; i++ {
		fmt.Printf("%c\n", kaisa(i))
	}

	// fmt.Printf("%v\n", len("abcdefg"))
}

func kaisa(c rune) rune {
	temp := c + 2
	if temp > 'Z' {
		temp -= 26
	}

	return temp
}
