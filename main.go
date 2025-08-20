package main

import (
	"fmt"

	"github.com/radifan9/minitask-w8-d2/internal/person"
	readfile "github.com/radifan9/minitask-w8-d2/internal/readFile"
)

func main() {
	fmt.Println("--- --- Minitask W8 D2 --- ---")

	// --- --- No 1 --- ---
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

	// --- --- No 2 --- ---
	user1 := person.NewPerson("Opet", "Acropolis", "08515519")
	user1.PrintData()
	user1.Greet("Ce")
	user1.UpdateName("Farid")
	user1.Greet("Ce")
}
