package pointersanderrors

import "fmt"

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
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

//Balance of Wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//String - formats Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
