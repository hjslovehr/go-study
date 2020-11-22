package main

import "fmt"

func main() {
	fmt.Printf("%v\n", "custom type demo")

	var pi rune = 960
	fmt.Printf("%v\n", pi)

	var bang byte = 33
	fmt.Printf("%v\n", bang)
	fmt.Printf("%c\n", bang)
	fmt.Printf("%T\n", bang)

	grade := 'A'
	fmt.Printf("%v, %c\n", grade, grade)

	var grade1 = 'B'
	fmt.Printf("%v, %c\n", grade1, grade1)

	var grade2 rune = 'C'
	fmt.Printf("%v, %c\n", grade2, grade2)

	message := "hello"
	temp := message[0]
	fmt.Printf("%c\n", temp)

	temp = 'D'
	fmt.Printf("%c\n", temp)

	// message[0] = 'A'	// 这句会报错 cannot assign to message[0]
}
