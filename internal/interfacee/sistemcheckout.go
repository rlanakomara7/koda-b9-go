package checkout

import (
	"errors"
	"fmt"
)

type Payment interface {
	Pay([]int) error
}

type Bank struct{}
type Online struct{}
type Fiktif struct {
	Total []int
}

func totalHarga(list []int) (int, error) {
	total := 0

	for _, harga := range list {
		if harga <= 0 {
			return 0, errors.New("harga harus lebih dari 0")
		}
		total += harga
	}

	return total, nil
}

func (Bank) Pay(list []int) error {
	total, err := totalHarga(list)
	if err != nil {
		return err
	}

	fmt.Println("Bayar melalui Bank:", total)
	return nil
}

func (Online) Pay(list []int) error {
	total, err := totalHarga(list)
	if err != nil {
		return err
	}

	fmt.Println("Bayar melalui Online:", total)
	return nil
}

func (f *Fiktif) Pay(list []int) error {
	total, err := totalHarga(list)
	if err != nil {
		return err
	}

	f.Total = append(f.Total, total)
	return nil
}

func Checkout(payment Payment, list []int) error {
	return payment.Pay(list)
}
