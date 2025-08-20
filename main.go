package main

import (
	"fmt"

	readfile "github.com/radifan9/minitask-w8-d2/internal/readFile"
)

func main() {
	fmt.Println("--- --- Minitask W8 D2 --- ---")

	text, err := readfile.MustReadFile("data/super-secret.text")
	// text, err := readfile.MustReadFile("data/super-secrett.text")
	// text, err := readfile.MustReadFile("data/")

	// Error handling (invalid path)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// Read file success
		fmt.Println(text)
	}

}
