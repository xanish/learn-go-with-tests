package _6_pointers_and_errors

type Wallet struct {
	// In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then it is private
	// outside the package it's defined in.
	balance int
}

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}
