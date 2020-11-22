package main

import (
	"fmt"
	"net/http"
)

func main() {
	var host = "localhost:6010"
	fmt.Println("Web start", host)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	err := http.ListenAndServe(host, nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
