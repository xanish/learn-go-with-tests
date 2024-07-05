package _6_pointers_and_errors

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds is an error returned when a wallet has insufficient funds for a withdrawal.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin represents a number of Bitcoins.
type Bitcoin int

// String formats the Bitcoin type for printing.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet represents a Bitcoin wallet with a balance.
type Wallet struct {
	balance Bitcoin
}

// Balance returns the current balance of the wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit adds the specified amount to the wallet's balance.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw subtracts the specified amount from the wallet's balance.
// Returns ErrInsufficientFunds if the amount exceeds the balance.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
