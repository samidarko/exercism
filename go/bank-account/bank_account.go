package account

import "sync"

// Account type here.
type Account struct {
	balance int64
	closed  bool
	sync.RWMutex
}

// Open and Account
func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{
		balance: amount,
	}
}

// Balance returns Account balance
func (a *Account) Balance() (int64, bool) {
	a.RLock()
	defer a.RUnlock()
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

// Deposit or withdraw amount
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	newBalance := a.balance + amount
	if amount < 0 && newBalance < 0 {
		return 0, false
	}
	a.balance = newBalance
	return a.balance, true
}

// Close an Account
func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	balance := a.balance
	a.balance = 0
	a.closed = true
	return balance, true
}
