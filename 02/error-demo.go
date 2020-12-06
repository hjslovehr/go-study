package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("---------- error demo ----------")
	test1()
}

func test1() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
