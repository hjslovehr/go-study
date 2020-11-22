package main

import "fmt"

func main() {
	fmt.Printf("%v\n", "range demo")

	arr := []int{1, 2, 3, 4, 5}

	for i, v := range arr {
		fmt.Printf("%v, %v\n", i, v)
	}

	fmt.Println()

	for _, v := range arr {
		fmt.Printf("%v\n", v)
	}

	fmt.Println()

	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
