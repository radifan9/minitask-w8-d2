package main

import (
	"fmt"

	"github.com/radifan9/minitask-w8-d2/internal/checkout"
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

	// --- --- No 3 --- ---

	amountNeedToPay1 := []int{1000, 2000, 3000}
	amountNeedToPay2 := []int{3000, 6000, 3000}

	bank1 := checkout.Bank{Name: "BCA"}
	result1, err1 := bank1.Pay(amountNeedToPay1)
	if err1 != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result1)
	}

	// Pembayaran fiktif
	fiktifPayment := &checkout.Fiktif{}
	result2, err2 := fiktifPayment.Pay(amountNeedToPay1)
	if err2 != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result2)
	}

	result3, err3 := fiktifPayment.Pay(amountNeedToPay2)
	if err3 != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result3)
	}

	fmt.Println("List pembayaran fiktif : ", fiktifPayment.Lists)
}
