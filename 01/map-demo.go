package main

import (
	"fmt"
)

func main() {
	fmt.Printf("========== map demo ==========\n")

	temp := map[string]int{
		"Earth": 15,
		"Mars":  65,
	}

	fmt.Printf("%v\n", temp)

	fmt.Println(temp["Earth"])
	fmt.Println(temp["Mars"])

	temp["Venus"] = -145

	fmt.Println("------------------------------------------------------------")

	moon := temp["Moon"]
	fmt.Println(moon)

	temp["Moon"] = -10

	for key, value := range temp {
		fmt.Println(key, value)
	}
}
