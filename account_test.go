package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	email := "test@example.com"
	password := "test_password"

	account := CreateAccount(email, password)

	assert.NotNil(t, account, "The account should not be nil")
	assert.Equal(t, email, account.Email, "The email should match")
	assert.Equal(t, password, account.Password, "The password should match")
}

func TestGetAccountByEmail(t *testing.T) {
	email := "test@example.com"
	password := "test_password"

	account := CreateAccount(email, password)

	foundAccount := GetAccountByEmail(email)

	assert.NotNil(t, foundAccount, "The found account should not be nil")
	assert.Equal(t, account.ID, foundAccount.ID, "The account IDs should match")
	assert.Equal(t, email, foundAccount.Email, "The email should match")
}
