package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	a.SetString("1234567890123456789", 10)
	fmt.Println(a)

	// const distance = 24000000000000000000	// 声明常量
	// fmt.Println(distance)	// 报错，原因是 distance 超过 int
}
