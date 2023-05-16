package main

import (
	"sync"
)

type Account struct {
	ID       int64
	Email    string
	Password string
}

var accountStore = struct {
	sync.RWMutex
	accounts map[int64]*Account
}{accounts: make(map[int64]*Account)}

func CreateAccount(email, password string) *Account {
	accountStore.Lock()
	defer accountStore.Unlock()

	id := int64(len(accountStore.accounts) + 1)
	account := &Account{
		ID:       id,
		Email:    email,
		Password: password,
	}

	accountStore.accounts[id] = account
	return account
}

func GetAccountByEmail(email string) *Account {
	accountStore.RLock()
	defer accountStore.RUnlock()

	for _, account := range accountStore.accounts {
		if account.Email == email {
			return account
		}
	}
	return nil
}
