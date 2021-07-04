package account

import "sync"

type Account struct {
	balance int64
	open    bool
	mu      sync.RWMutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{
		balance: initialDeposit,
		open:    true,
	}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.open {
		return
	}

	a.open = false

	return a.balance, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if !a.open {
		return
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.open {
		return
	}

	if a.balance + amount < 0 {
		return
	}

	a.balance += amount

	return a.balance, true
}
