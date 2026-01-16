package interfaces

import "fmt"

type Card struct {
	Balance int
}

func (c *Card) Pay(price int) error {
	if c.Balance < price {
		return fmt.Errorf("not enough money")
	}
	c.Balance -= price
	return nil
}

type Wallet struct {
	Balance int
}

func (w *Wallet) Pay(price int) error {
	if w.Balance < price {
		return fmt.Errorf("not enough money")
	}
	w.Balance -= price
	return nil
}

type ApplePay struct {
	Balance int
}

func (ap *ApplePay) Pay(price int) error {
	if ap.Balance < price {
		return fmt.Errorf("not enough money")
	}
	ap.Balance -= price
	return nil
}
