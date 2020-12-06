package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("---------- file demo ----------")

	err := proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	_, errerr = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		f.Close()
		return err
	}

	_, err = fmt.Fprintln(f, "Do not just check errors, handle them gracefully.")
	f.Close()
	return err
}
