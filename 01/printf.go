package main

import (
	"fmt"
)

func main() {
	hello("Go")
}

func hello(string name) {
	fmt.Printf("Hello, %v", name)
}
