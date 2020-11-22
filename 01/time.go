package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("========== time demo ==========")

	now := time.Now()
	nowStr := "time.Now()"

	fmt.Printf("%v: %v\n", nowStr, now)
	fmt.Printf("%v.Format(\"2006-01-02 15:04:05\"): %v\n", nowStr, now.Format("2006-01-02 15:04:05"))

	fmt.Printf("%v.Year(): %v\n", nowStr, now.Year())
	fmt.Printf("%v.Month(): %v\n", nowStr, int(now.Month()))
	fmt.Printf("%v.Day(): %v\n", nowStr, now.Day())

	year, month, day := time.Now().Date()
	fmt.Println("Year:", year, "Month:", month, "Day", day)      // For example 2009 November 10
	fmt.Println("Year:", year, "Month:", int(month), "Day", day) // For example 2009 11 10
}
