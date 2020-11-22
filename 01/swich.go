package main

import "fmt"

func main() {
	var value = "csharp"
	switch value {
	case "go":
		fmt.Printf("is go")
	case "csharp":
		fmt.Printf("is csharp")
	default:
		fmt.Printf("no match")
	}
}
