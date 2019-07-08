package pointersanderrors

import (
	"errors"
	"fmt"
)

//Bitcoin - new
type Bitcoin int

//Stringer - new
type Stringer interface {
	String() string
}

//Wallet - struct
type Wallet struct {
	balance Bitcoin
}

//Deposit to Wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//ErrInsufficientFunds for errors
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Withdraw from Wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

//Balance of Wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//String - formats Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
