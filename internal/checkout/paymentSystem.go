package checkout

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount []int) (string, error)
}

type Bank struct {
	Name string
}

type Online struct {
	Name string
}

type Fiktif struct {
	Lists []int
}

func (b *Bank) Pay(amount []int) (string, error) {

	total := 0
	for _, value := range amount {
		total = total + value
	}

	return fmt.Sprintf("Pembayaran %d berhasil via Bank %s\n", total, b.Name), nil
}

func (o *Online) Pay(amount []int) (string, error) {
	total := 0
	for _, value := range amount {
		total += value
	}
	return fmt.Sprintf("Pembayaran %d berhasil via Online %s\n", total, o.Name), nil
}

func (f *Fiktif) Pay(amount []int) (string, error) {
	total := 0
	for _, value := range amount {
		total = total + value
	}

	if total <= 0 {
		return "", errors.New("total pembayaran kurang dari 0")
	}

	f.Lists = append(f.Lists, total)
	return fmt.Sprintf("Pembayaran %d berhasil via Fiktif\n", total), nil
}
