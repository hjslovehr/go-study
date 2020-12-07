package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpGet(url string) (string, string, error) {
	resp, err := http.Get("http://www.baidu.com/")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return resp.Status, string(bodyContent), nil
}

func getTest() {
	status, body, err := httpGet("https://www.baidu.com/")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("status: ", status)
	fmt.Println("body: ", body)
}

func main() {
	fmt.Println("---------- http request demo ----------")
	getTest()
}
