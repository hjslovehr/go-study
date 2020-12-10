package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func test() {
	resp, err := http.Get("http://www.baidu.com/")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bodyContent))
}

func main() {
	fmt.Println("---------- http demo ----------")
	test()
}
