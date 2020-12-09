package main

import (
	"errors"
	"fmt"
)

func test() error {
	return errors.New("this is a test error")
}

func main() {
	fmt.Println("---------- panic demo ----------")

	err := test()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
