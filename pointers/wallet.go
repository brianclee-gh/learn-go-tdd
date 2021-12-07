package main

import "errors"

var ErrInsufficientFunds = errors.New("insufficient funds")

type Bitcoin int

type Wallet struct {
	total Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.total
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.total += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.total {
		return ErrInsufficientFunds
	}
	w.total -= amount
	return nil
}
