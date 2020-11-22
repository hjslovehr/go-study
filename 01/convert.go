package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("========== convert demo ==========")

	fmt.Println("int to string")

	a := strconv.Itoa(1)
	fmt.Printf("数字转字符串 1 => str : %v\n", a)

	{
		val := "123"
		// b, err := strconv.Atoi(val)	// 这个是 Atoi是ParseInt(s, 10, 0) 的简写
		b, err := strconv.ParseInt(val, 10, 0)
		if err != nil {
			fmt.Println("字符串转整数数字失败")
		} else {
			fmt.Printf("%v => Int: %v \n", val, b)
		}
	}

	{
		val := "123a"
		b, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("字符串转整数数字失败")
		} else {
			fmt.Printf("%v => Int: %v \n", val, b)
		}
	}

	{
		val := "3.1415926"
		b, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Println("字符串转浮点数字失败")
		} else {
			fmt.Printf("%v => Float32: %v \n", val, b)
		}
	}

	{
		val := "3.1415926abc"
		b, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Println("字符串转浮点数字失败")
		} else {
			fmt.Printf("%v => Float32: %v \n", val, b)
		}
	}
}
