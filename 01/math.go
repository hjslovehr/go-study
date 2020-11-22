package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("========== math demo ==========")
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		v := rand.Intn(100)
		fmt.Printf("%v: %v\n", time.Now().Format("2006-01-02 15:04:05"), v)
	}
}
