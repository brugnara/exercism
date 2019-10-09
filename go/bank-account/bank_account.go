package account

import (
	"sync"
)

// Account a
type Account struct {
	closed bool
	money  int64
	mutex  *sync.Mutex
}

// Open o
func Open(initialDeposit int64) (account *Account) {
	if initialDeposit < 0 {
		return nil
	}
	account = new(Account)
	account.closed = false
	account.money = initialDeposit
	account.mutex = new(sync.Mutex)
	return
}

// Balance will return balance
func (a *Account) Balance() (amount int64, ok bool) {
	if a.closed {
		return
	}
	return a.money, true
}

// Close closes the account. We must ensure this func will be executed with
// an eye to concurrency so we use Mutex here. To be sure to Unlock() the mutex
// I use defer. This allow us to simplify the code without using waitGroups
func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.closed {
		return
	}
	a.closed = true
	payout = a.money
	a.money = 0
	ok = true
	return
}

// Deposit deposit some money, here again we are using Mutex. The whole func
// must be locked in order to have right returing values for balance and ok
func (a *Account) Deposit(money int64) (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	// cannot deposit if account is closed or the new balance will be negative
	if a.closed || a.money+money < 0 {
		return
	}

	// depositing
	a.money += money

	balance = a.money
	ok = true
	return
}
