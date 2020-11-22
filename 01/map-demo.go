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

	temp["Venus"] = 464

	fmt.Println("------------------------------------------------------------")

	moon := temp["Moon"]
	fmt.Println(moon)

	temp["Moon"] = -10

	for key, value := range temp {
		fmt.Println(key, value)
	}

	if abc, ok := temp["abc"]; ok {
		fmt.Println(abc)
		fmt.Println("abc is exists")
	} else {
		fmt.Println("abc is not exists")
	}
}
