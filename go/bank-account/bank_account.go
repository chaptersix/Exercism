package account

import "sync"

//Account represents an account
type Account struct {
	balance int64
	closed  bool
	m       sync.Mutex
}

// Open initializes and Account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	a := Account{balance: initialDeposit}
	return &a
}

// Close closes the Account
func (a *Account) Close() (payout int64, ok bool) {
	a.m.Lock()
	defer a.m.Unlock()
	if a.closed {
		return 0, false
	}
	a.closed = true
	payout = a.balance
	a.balance = 0
	return payout, true
}

//Balance return the balance of the Account
func (a *Account) Balance() (balance int64, ok bool) {
	a.m.Lock()
	defer a.m.Unlock()
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

//Deposit adds to the Account balance
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.m.Lock()
	defer a.m.Unlock()
	if a.closed || amount+a.balance < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
