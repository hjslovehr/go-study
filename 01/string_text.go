package main

import "fmt"

func main() {
	fmt.Println("string and text demo")
	s := "abc"
	fmt.Printf("%v", s) // 输出 abc

	var s1 string
	fmt.Println(s1) // 什么都不输出

	var s2 = "price"
	fmt.Println(s2) // 输出 price

	fmt.Println("abc\n123")

	fmt.Println(`aaa\nbbb`) // 输出 aaa\nbbb

	fmt.Println(`
	abc
	def
	ghi
	`) // 原封不动的输出字符串，并保留换行，空格等空白字符

	fmt.Println(`c:\go`)
}
